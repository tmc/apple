#include <bit>
#include <cmath>
#include <cstdint>
#include <cstdio>
#include <limits>
#include <span>

#include "jaccl/group.h"
#include "jaccl/reduction_ops.h"

int main() {
  jaccl::float16_t half_nan = std::numeric_limits<float>::quiet_NaN();
  jaccl::float16_t half_large = 65520.0f;
  jaccl::float16_t half_one = 1.0f;
  jaccl::float16_t half_input = std::numeric_limits<float>::quiet_NaN();
  jaccl::SumOp<jaccl::float16_t>{}(&half_input, &half_one, 1);

  jaccl::bfloat16_t bfloat_nan = std::numeric_limits<float>::quiet_NaN();

  std::printf(
      "%04x %04x %04x %04x\n",
      std::bit_cast<std::uint16_t>(half_nan),
      std::bit_cast<std::uint16_t>(half_large),
      std::bit_cast<std::uint16_t>(half_one),
      bfloat_nan.bits_);

  auto print_value = [](const auto& value) {
    for (std::byte byte : std::as_bytes(std::span{&value, 1})) {
      std::printf("%02x", std::to_integer<unsigned char>(byte));
    }
  };
  auto emit = [&](auto a, auto b) {
    auto sum = a;
    auto max = a;
    auto min = a;
    using T = decltype(a);
    jaccl::SumOp<T>{}(&b, &sum, 1);
    jaccl::MaxOp<T>{}(&b, &max, 1);
    jaccl::MinOp<T>{}(&b, &min, 1);
    print_value(sum);
    std::printf(" ");
    print_value(max);
    std::printf(" ");
    print_value(min);
    std::printf("\n");
  };
  emit(true, false);
  emit(int8_t{-2}, int8_t{5});
  emit(int16_t{-2}, int16_t{5});
  emit(int32_t{-2}, int32_t{5});
  emit(int64_t{-2}, int64_t{5});
  emit(uint8_t{2}, uint8_t{5});
  emit(uint16_t{2}, uint16_t{5});
  emit(uint32_t{2}, uint32_t{5});
  emit(uint64_t{2}, uint64_t{5});
  emit(jaccl::float16_t{2.25f}, jaccl::float16_t{1.5f});
  emit(jaccl::bfloat16_t{2.25f}, jaccl::bfloat16_t{1.5f});
  emit(2.25f, 1.5f);
  emit(2.25, 1.5);
  emit(jaccl::complex64_t{2.0f, 5.0f}, jaccl::complex64_t{1.0f, 8.0f});

  float rank_order = 1e20f;
  float rank_one = -1e20f;
  float rank_two = 1.0f;
  jaccl::SumOp<float>{}(&rank_one, &rank_order, 1);
  jaccl::SumOp<float>{}(&rank_two, &rank_order, 1);
  print_value(rank_order);
  std::printf("\n");

  constexpr bool bool_values[] = {false, true};
  for (bool a : bool_values) {
    for (bool b : bool_values) {
      emit(a, b);
    }
  }

  constexpr int8_t int8_values[] = {-64, -1, 0, 1, 63};
  for (int8_t a : int8_values) {
    for (int8_t b : int8_values) {
      emit(a, b);
    }
  }
  constexpr int16_t int16_values[] = {-16384, -1, 0, 1, 16383};
  for (int16_t a : int16_values) {
    for (int16_t b : int16_values) {
      emit(a, b);
    }
  }
  constexpr int32_t int32_values[] = {-(1 << 29), -1, 0, 1, (1 << 29) - 1};
  for (int32_t a : int32_values) {
    for (int32_t b : int32_values) {
      emit(a, b);
    }
  }
  constexpr int64_t int64_values[] = {-(int64_t{1} << 61), -1, 0, 1, (int64_t{1} << 61) - 1};
  for (int64_t a : int64_values) {
    for (int64_t b : int64_values) {
      emit(a, b);
    }
  }
  constexpr uint8_t uint8_values[] = {0, 1, 2, 63, 127};
  for (uint8_t a : uint8_values) {
    for (uint8_t b : uint8_values) {
      emit(a, b);
    }
  }
  constexpr uint16_t uint16_values[] = {0, 1, 2, 0x3fff, 0x7fff};
  for (uint16_t a : uint16_values) {
    for (uint16_t b : uint16_values) {
      emit(a, b);
    }
  }
  constexpr uint32_t uint32_values[] = {0, 1, 2, 0x3fffffff, 0x7fffffff};
  for (uint32_t a : uint32_values) {
    for (uint32_t b : uint32_values) {
      emit(a, b);
    }
  }
  constexpr uint64_t uint64_values[] = {0, 1, 2, UINT64_MAX / 4, UINT64_MAX / 2};
  for (uint64_t a : uint64_values) {
    for (uint64_t b : uint64_values) {
      emit(a, b);
    }
  }

  constexpr std::uint16_t half_values[] = {
      0x0000, 0x8000, 0x0001, 0x03ff, 0x0400, 0x3555, 0x3c00,
      0x3c01, 0x7bff, 0x7c00, 0xfc00, 0x7d00, 0x7e00, 0xfe00,
  };
  for (std::uint16_t a_bits : half_values) {
    for (std::uint16_t b_bits : half_values) {
      emit(
          std::bit_cast<jaccl::float16_t>(a_bits),
          std::bit_cast<jaccl::float16_t>(b_bits));
    }
  }

  constexpr std::uint16_t bfloat_values[] = {
      0x0000, 0x8000, 0x0001, 0x3f80, 0x4000, 0x7f7f, 0x7f80, 0xff80,
      0x7f81, 0x7fc0,
  };
  for (std::uint16_t a_bits : bfloat_values) {
    for (std::uint16_t b_bits : bfloat_values) {
      jaccl::bfloat16_t a{0.0f};
      jaccl::bfloat16_t b{0.0f};
      a.bits_ = a_bits;
      b.bits_ = b_bits;
      emit(a, b);
    }
  }

  constexpr std::uint32_t float_values[] = {
      0x00000000, 0x80000000, 0x3f800000, 0xbf800000,
      0x7f800000, 0xff800000, 0x7f800001, 0x7fc00000, 0xffc00000,
  };
  for (std::uint32_t a_bits : float_values) {
    for (std::uint32_t b_bits : float_values) {
      emit(std::bit_cast<float>(a_bits), std::bit_cast<float>(b_bits));
    }
  }

  constexpr std::uint64_t double_values[] = {
      0x0000000000000000, 0x8000000000000000, 0x3ff0000000000000,
      0xbff0000000000000, 0x7ff0000000000000, 0xfff0000000000000,
      0x7ff0000000000001, 0x7ff8000000000000, 0xfff8000000000000,
  };
  for (std::uint64_t a_bits : double_values) {
    for (std::uint64_t b_bits : double_values) {
      emit(std::bit_cast<double>(a_bits), std::bit_cast<double>(b_bits));
    }
  }

  constexpr std::uint32_t complex_values[][2] = {
      {0x00000000, 0x80000000},
      {0x3f800000, 0xbf800000},
      {0x7f800001, 0xff800001},
      {0x7fc00000, 0xffc00000},
      {0xffc00000, 0x7fc00000},
  };
  for (const auto& a : complex_values) {
    for (const auto& b : complex_values) {
      emit(
          jaccl::complex64_t{
              std::bit_cast<float>(a[0]), std::bit_cast<float>(a[1])},
          jaccl::complex64_t{
              std::bit_cast<float>(b[0]), std::bit_cast<float>(b[1])});
    }
  }
}
