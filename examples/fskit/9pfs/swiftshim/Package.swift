// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "NinePFSShim",
    platforms: [
        .macOS("15.4")
    ],
    products: [
        .library(
            name: "NinePFSShim",
            type: .dynamic,
            targets: ["NinePFSShim"]
        )
    ],
    targets: [
        .target(
            name: "NinePFSShimObjC",
            publicHeadersPath: "include"
        ),
        .target(
            name: "NinePFSShim",
            dependencies: ["NinePFSShimObjC"]
        )
    ]
)
