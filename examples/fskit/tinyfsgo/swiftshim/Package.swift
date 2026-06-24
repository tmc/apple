// swift-tools-version: 6.0

import PackageDescription

let package = Package(
    name: "TinyFSShim",
    platforms: [
        .macOS("15.4")
    ],
    products: [
        .library(
            name: "TinyFSShim",
            type: .dynamic,
            targets: ["TinyFSShim"]
        )
    ],
    targets: [
        .target(
            name: "TinyFSShimObjC",
            publicHeadersPath: "include"
        ),
        .target(
            name: "TinyFSShim",
            dependencies: ["TinyFSShimObjC"]
        )
    ]
)
