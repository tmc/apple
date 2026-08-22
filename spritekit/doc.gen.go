// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

// Package spritekit provides Go bindings for the SpriteKit framework.
//
// Add high-performance 2D content with smooth animations to your app, or
// create a game with a high-level set of 2D game-based tools.
//
// SpriteKit is a general-purpose framework for drawing shapes, particles,
// text, images, and video in two dimensions. It leverages Metal to achieve
// high-performance rendering, while offering a simple programming interface
// to make it easy to create games and other graphics-intensive apps. Using a
// rich set of animations and physics behaviors, you can quickly add life to
// your visual elements and gracefully transition between screens.
//
// # Essentials
//
//   - [Drawing SpriteKit Content in a View]: Display visual content using SpriteKit.
//   - [SKScene]: An object that organizes all of the active SpriteKit content. ([SKSceneScaleMode], [SKSceneDelegate])
//   - [Nodes for Scene Building]: Define the appearance or layout of scene content. ([SKNode], [SKCameraNode], [SKReferenceNode], [SKSpriteNode], [SKShapeNode])
//
// # Scene Renderers
//
//   - [Choosing a SpriteKit Scene Renderer]: Compare the different ways to display a SpriteKit scene.
//   - [SKView]: A view subclass that renders a SpriteKit scene. ([SKTransition], [SKViewDelegate])
//   - [SKRenderer]: An object that renders a scene into a custom Metal rendering pipeline and drives the scene update cycle.
//
// # Textures
//
//   - [Maximizing Texture Performance]: Speed up image display and enable more images to be displayed at one time.
//   - [SKTexture]: An image, decoded on the GPU, that can be used to render various SpriteKit objects. ([SKTextureFilteringMode])
//   - [SKTextureAtlas]: A collection of textures optimized for storage and drawing performance.
//   - [SKMutableTexture]: A texture whose contents can be dynamically updated.
//
// # Animation
//
//   - [Getting Started with Actions]: Create, configure, and run actions in SpriteKit.
//   - [SKAction]: An object that is run by a node to change its structure or content. ([SKActionTimingMode], [SKActionTimingFunction])
//
// # Constraints
//
//   - [SKConstraint]: A specification for constraining a node’s position or rotation.
//   - [SKReachConstraints]: A specification of the degree of freedom when solving inverse kinematics.
//
// # Mathematical Tools
//
//   - [SKKeyframeSequence]: An object that performs interpolation between values specified at different times (keyframes). ([SKInterpolationMode], [SKRepeatMode])
//   - [SKRange]: A definition of a range of floating-point values.
//   - [SKRegion]: The definition of an arbitrary area.
//
// # Physics Simulation
//
//   - [Getting Started with Physics]: Simulate gravity, acceleration, collision detection, or joints.
//   - [SKPhysicsWorld]: The driver of the physics engine in a scene; it exposes the ability for you to configure and query the physics system.
//   - [SKPhysicsBody]: An object that adds physics simulation to a node.
//   - [SKPhysicsContact]: A description of the contact between two physics bodies.
//   - [SKPhysicsContactDelegate]: Methods your app can implement to respond when physics bodies come into contact.
//   - [SKFieldNode]: A node that applies physics effects to nearby nodes. ([SKFieldForceEvaluator])
//
// # Physics Joints
//
//   - [Working with Inverse Kinematics]: Gain fine-tuned control of objects that are connected by joints.
//   - [SKPhysicsJoint]: The abstract superclass for objects that connect physics bodies.
//   - [SKPhysicsJointFixed]: A joint that fuses two physics bodies together at a reference point.
//   - [SKPhysicsJointLimit]: A joint that imposes a maximum distance between two physics bodies, as if they were connected by a rope.
//   - [SKPhysicsJointPin]: A joint that pins together two physics bodies, allowing independent rotation.
//   - [SKPhysicsJointSliding]: A joint that allows two physics bodies to slide along an axis.
//   - [SKPhysicsJointSpring]: A joint that simulates a spring connecting two physics bodies.
//
// # Tiling
//
//   - [SKTileMapNode]: A two-dimensional array of images.
//   - [SKTileDefinition]: A single tile that can be repeated in a tile map. ([SKTileDefinitionRotation])
//   - [SKTileGroup]: A set of tiles that collectively define one type of terrain.
//   - [SKTileGroupRule]: Rules that describe how various tiles should be placed in a map. ([SKTileAdjacencyMask])
//   - [SKTileSet]: A container for related tile groups. ([SKTileSetType])
//
// # Shaders
//
//   - [SKShader]: An object that allows you to apply a custom fragment shader.
//   - [SKAttribute]: A specification for dynamic per-node data used with a custom shader. ([SKAttributeType])
//   - [SKAttributeValue]: A container for dynamic shader data associated with a node.
//   - [SKUniform]: A container for uniform shader data. ([SKUniformType])
//
// # Warping
//
//   - [SKWarpGeometry]: A definition for a deformation of nodes that conform to .
//   - [SKWarpGeometryGrid]: A definition for a grid-based deformation of nodes that conform to .
//   - [SKWarpable]: A protocol for objects that can be warped and animated by an .//
//
// # Key Types
//
//   - [SKNode] - The base class of all SpriteKit nodes.
//   - [SKEmitterNode] - A source of various particle effects.
//   - [SKPhysicsBody] - An object that adds physics simulation to a node.
//   - [SKShapeNode] - A mathematical shape that can be stroked or filled.
//   - [SKSpriteNode] - An image or solid color.
//   - [SKScene] - An object that organizes all of the active SpriteKit content.
//   - [SKFieldNode] - A node that applies physics effects to nearby nodes.
//   - [SKTileMapNode] - A two-dimensional array of images.
//   - [SKView] - A view subclass that renders a SpriteKit scene.
//   - [SKLabelNode] - A graphical element that draws text.
//
// [Choosing a SpriteKit Scene Renderer]: https://developer.apple.com/documentation/spritekit/choosing-a-spritekit-scene-renderer
// [Drawing SpriteKit Content in a View]: https://developer.apple.com/documentation/spritekit/drawing-spritekit-content-in-a-view
// [Getting Started with Actions]: https://developer.apple.com/documentation/spritekit/getting-started-with-actions
// [Getting Started with Physics]: https://developer.apple.com/documentation/spritekit/getting-started-with-physics
// [Maximizing Texture Performance]: https://developer.apple.com/documentation/spritekit/maximizing-texture-performance
// [Nodes for Scene Building]: https://developer.apple.com/documentation/spritekit/nodes-for-scene-building
// [Working with Inverse Kinematics]: https://developer.apple.com/documentation/spritekit/working-with-inverse-kinematics
package spritekit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the SpriteKit library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/SpriteKit.framework/SpriteKit",
	"/usr/lib/libSpriteKit.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: SpriteKit: failed to load framework from any known path\n")
	}
}
