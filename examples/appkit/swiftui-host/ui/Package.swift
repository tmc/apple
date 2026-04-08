// swift-tools-version: 5.9
import PackageDescription

let package = Package(
    name: "ExampleAppUI",
    platforms: [
        .macOS(.v13),
    ],
    products: [
        .library(name: "ExampleAppUI", type: .dynamic, targets: ["ExampleAppUI"]),
    ],
    targets: [
        .target(name: "ExampleAppUI"),
    ]
)
