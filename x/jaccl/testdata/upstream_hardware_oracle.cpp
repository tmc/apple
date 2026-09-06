#include <CommonCrypto/CommonDigest.h>

#include <array>
#include <cstdint>
#include <cstdio>
#include <exception>
#include <memory>
#include <span>
#include <stdexcept>
#include <string>
#include <type_traits>
#include <vector>

#include "jaccl/jaccl.h"
#include "jaccl/types.h"

namespace {

std::string sha256(std::span<const std::byte> data) {
  std::array<unsigned char, CC_SHA256_DIGEST_LENGTH> digest;
  CC_SHA256(data.data(), static_cast<CC_LONG>(data.size()), digest.data());
  std::string result;
  result.reserve(digest.size() * 2);
  for (unsigned char byte : digest) {
    char text[3];
    std::snprintf(text, sizeof(text), "%02x", byte);
    result += text;
  }
  return result;
}

template <typename T, std::size_t N>
std::span<const std::byte> bytes(const std::array<T, N>& values) {
  return std::as_bytes(std::span{values});
}

template <typename T>
bool equal(T got, T want) {
  if constexpr (
      std::is_same_v<T, jaccl::float16_t> ||
      std::is_same_v<T, jaccl::bfloat16_t>) {
    return static_cast<float>(got) == static_cast<float>(want);
  }
  return got == want;
}

template <typename T>
void append_bytes(std::vector<std::byte>& dst, const T& value) {
  const auto value_bytes = std::as_bytes(std::span<const T, 1>(&value, 1));
  dst.insert(dst.end(), value_bytes.begin(), value_bytes.end());
}

template <typename T>
void all_reduce_receipt(
    jaccl::Group& group,
    int rank,
    int size,
    jaccl::Dtype dtype,
    std::vector<std::byte>& receipt) {
  const T input = static_cast<T>(rank + 1);
  T sum = static_cast<T>(0);
  T maximum = static_cast<T>(0);
  T minimum = static_cast<T>(0);
  group.all_sum(&input, &sum, sizeof(input), dtype);
  group.all_max(&input, &maximum, sizeof(input), dtype);
  group.all_min(&input, &minimum, sizeof(input), dtype);

  const T want_sum = static_cast<T>(size * (size + 1) / 2);
  const T want_maximum = static_cast<T>(size);
  const T want_minimum = static_cast<T>(1);
  if (!equal(sum, want_sum) || !equal(maximum, want_maximum) ||
      !equal(minimum, want_minimum)) {
    throw std::runtime_error("upstream JACCL all-dtype all-reduce differs");
  }
  append_bytes(receipt, sum);
  append_bytes(receipt, maximum);
  append_bytes(receipt, minimum);
}

} // namespace

int main() {
  try {
    std::fprintf(stderr, "phase=init start\n");
    std::shared_ptr<jaccl::Group> group = jaccl::init(/* strict= */ true);
    const int rank = group->rank();
    const int size = group->size();
    std::fprintf(stderr, "phase=init complete rank=%d size=%d\n", rank, size);
    if (size < 2 || rank < 0 || rank >= size) {
      std::fprintf(stderr, "upstream JACCL hardware oracle requires valid ranks\n");
      return 2;
    }

    const std::array<std::byte, 31> payload = {
        std::byte{'j'}, std::byte{'a'}, std::byte{'c'}, std::byte{'c'},
        std::byte{'l'}, std::byte{' '}, std::byte{'h'}, std::byte{'a'},
        std::byte{'r'}, std::byte{'d'}, std::byte{'w'}, std::byte{'a'},
        std::byte{'r'}, std::byte{'e'}, std::byte{' '}, std::byte{'e'},
        std::byte{'v'}, std::byte{'i'}, std::byte{'d'}, std::byte{'e'},
        std::byte{'n'}, std::byte{'c'}, std::byte{'e'}, std::byte{' '},
        std::byte{'p'}, std::byte{'a'}, std::byte{'y'}, std::byte{'l'},
        std::byte{'o'}, std::byte{'a'}, std::byte{'d'},
    };
    if (rank == 0) {
      std::fprintf(stderr, "phase=p2p send\n");
      group->send(payload.data(), payload.size(), 1);
    } else if (rank == 1) {
      std::fprintf(stderr, "phase=p2p recv\n");
      std::array<std::byte, 31> got;
      group->recv(got.data(), got.size(), 0);
      if (got != payload) {
        std::fprintf(stderr, "upstream JACCL payload differs\n");
        return 3;
      }
    }

    std::fprintf(stderr, "phase=all-gather\n");
    const std::array<std::byte, 4> gather_input = {
        static_cast<std::byte>(rank), static_cast<std::byte>(rank + 10),
        static_cast<std::byte>(rank + 20), static_cast<std::byte>(rank + 30),
    };
    std::vector<std::byte> gathered(size * gather_input.size());
    group->all_gather(gather_input.data(), gathered.data(), gather_input.size());
    std::vector<std::byte> want_gathered;
    want_gathered.reserve(gathered.size());
    for (int peer = 0; peer < size; peer++) {
      want_gathered.insert(
          want_gathered.end(),
          {static_cast<std::byte>(peer), static_cast<std::byte>(peer + 10),
           static_cast<std::byte>(peer + 20), static_cast<std::byte>(peer + 30)});
    }
    if (gathered != want_gathered) {
      std::fprintf(stderr, "upstream JACCL all-gather differs\n");
      return 4;
    }

    std::fprintf(stderr, "phase=all-reduce-fixed\n");
    const std::array<std::int32_t, 2> reduce_input = {rank + 1, 5 - rank};
    std::array<std::int32_t, 2> reduced_sum;
    std::array<std::int32_t, 2> reduced_max;
    std::array<std::int32_t, 2> reduced_min;
    group->all_sum(
        reduce_input.data(), reduced_sum.data(), sizeof(reduce_input), jaccl::Int32);
    group->all_max(
        reduce_input.data(), reduced_max.data(), sizeof(reduce_input), jaccl::Int32);
    group->all_min(
        reduce_input.data(), reduced_min.data(), sizeof(reduce_input), jaccl::Int32);
    const std::array<std::int32_t, 2> want_sum = {
        size * (size + 1) / 2, 5 * size - size * (size - 1) / 2};
    const std::array<std::int32_t, 2> want_max = {size, 5};
    const std::array<std::int32_t, 2> want_min = {1, 6 - size};
    if (reduced_sum != want_sum || reduced_max != want_max || reduced_min != want_min) {
      std::fprintf(stderr, "upstream JACCL all-reduce differs\n");
      return 5;
    }

    const std::array<float, 2> float_input = {
        static_cast<float>(rank + 1), static_cast<float>(10 - rank)};
    std::array<float, 2> float_sum;
    std::array<float, 2> float_max;
    std::array<float, 2> float_min;
    group->all_sum(
        float_input.data(), float_sum.data(), sizeof(float_input), jaccl::Float32);
    group->all_max(
        float_input.data(), float_max.data(), sizeof(float_input), jaccl::Float32);
    group->all_min(
        float_input.data(), float_min.data(), sizeof(float_input), jaccl::Float32);
    const std::array<float, 2> want_float_sum = {
        static_cast<float>(size * (size + 1) / 2),
        static_cast<float>(10 * size - size * (size - 1) / 2)};
    const std::array<float, 2> want_float_max = {static_cast<float>(size), 10.0f};
    const std::array<float, 2> want_float_min = {1.0f, static_cast<float>(11 - size)};
    if (float_sum != want_float_sum || float_max != want_float_max || float_min != want_float_min) {
      std::fprintf(stderr, "upstream JACCL float32 all-reduce differs\n");
      return 6;
    }

    std::fprintf(stderr, "phase=all-reduce-all-dtypes\n");
    std::vector<std::byte> all_dtype_receipt;
    all_dtype_receipt.reserve(3 * (1 + 1 + 2 + 4 + 8 + 1 + 2 + 4 + 8 + 2 + 2 + 4 + 8 + 8));
    all_reduce_receipt<bool>(*group, rank, size, jaccl::Bool, all_dtype_receipt);
    all_reduce_receipt<std::int8_t>(*group, rank, size, jaccl::Int8, all_dtype_receipt);
    all_reduce_receipt<std::int16_t>(*group, rank, size, jaccl::Int16, all_dtype_receipt);
    all_reduce_receipt<std::int32_t>(*group, rank, size, jaccl::Int32, all_dtype_receipt);
    all_reduce_receipt<std::int64_t>(*group, rank, size, jaccl::Int64, all_dtype_receipt);
    all_reduce_receipt<std::uint8_t>(*group, rank, size, jaccl::UInt8, all_dtype_receipt);
    all_reduce_receipt<std::uint16_t>(*group, rank, size, jaccl::UInt16, all_dtype_receipt);
    all_reduce_receipt<std::uint32_t>(*group, rank, size, jaccl::UInt32, all_dtype_receipt);
    all_reduce_receipt<std::uint64_t>(*group, rank, size, jaccl::UInt64, all_dtype_receipt);
    all_reduce_receipt<jaccl::float16_t>(*group, rank, size, jaccl::Float16, all_dtype_receipt);
    all_reduce_receipt<jaccl::bfloat16_t>(*group, rank, size, jaccl::BFloat16, all_dtype_receipt);
    all_reduce_receipt<float>(*group, rank, size, jaccl::Float32, all_dtype_receipt);
    all_reduce_receipt<double>(*group, rank, size, jaccl::Float64, all_dtype_receipt);
    all_reduce_receipt<jaccl::complex64_t>(*group, rank, size, jaccl::Complex64, all_dtype_receipt);

    std::printf(
        "implementation=upstream-jaccl rank=%d size=%d payload_sha256=%s "
        "all_gather_sha256=%s all_reduce_int32_sum_sha256=%s "
        "all_reduce_int32_max_sha256=%s all_reduce_int32_min_sha256=%s "
        "all_reduce_float32_sum_sha256=%s all_reduce_float32_max_sha256=%s "
        "all_reduce_float32_min_sha256=%s all_reduce_all_dtypes_sha256=%s\n",
        rank,
        size,
        sha256(std::span{payload}).c_str(),
        sha256(std::span{gathered}).c_str(),
        sha256(bytes(reduced_sum)).c_str(),
        sha256(bytes(reduced_max)).c_str(),
        sha256(bytes(reduced_min)).c_str(),
        sha256(bytes(float_sum)).c_str(),
        sha256(bytes(float_max)).c_str(),
        sha256(bytes(float_min)).c_str(),
        sha256(std::span{all_dtype_receipt}).c_str());
    return 0;
  } catch (const std::exception& err) {
    std::fprintf(stderr, "upstream JACCL hardware oracle: %s\n", err.what());
    return 1;
  }
}
