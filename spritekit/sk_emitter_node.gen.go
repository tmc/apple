// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKEmitterNode] class.
var (
	_SKEmitterNodeClass     SKEmitterNodeClass
	_SKEmitterNodeClassOnce sync.Once
)

func getSKEmitterNodeClass() SKEmitterNodeClass {
	_SKEmitterNodeClassOnce.Do(func() {
		_SKEmitterNodeClass = SKEmitterNodeClass{class: objc.GetClass("SKEmitterNode")}
	})
	return _SKEmitterNodeClass
}

// GetSKEmitterNodeClass returns the class object for SKEmitterNode.
func GetSKEmitterNodeClass() SKEmitterNodeClass {
	return getSKEmitterNodeClass()
}

type SKEmitterNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKEmitterNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKEmitterNodeClass) Alloc() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A source of various particle effects.
//
// # Overview
//
// A [SKEmitterNode] object is a node that automatically creates and renders
// small particle sprites. Particles are privately owned by [SpriteKit]—your
// game cannot access the generated sprites. For example, you cannot add
// physics shapes to particles. Emitter nodes are often used to create smoke,
// fire, sparks, and other particle effects. A particle is similar to an
// [SKSpriteNode] object; it renders a textured or untextured image that is
// sized, colorized, and blended into the scene. However, particles differ
// from sprites in two important ways:
//
// - A particle’s texture is always stretched uniformly. - Particles are not
// represented by objects in SpriteKit. This means you cannot perform
// node-related tasks on particles, nor can you associate physics bodies with
// particles to make them interact with other content. Although there is no
// visible class representing particles added by the emitter node, you can
// think of a particle as having properties like any other object.
//
// Particles are purely visual objects, and their behavior is entirely defined
// by the emitter node that created them. The emitter node contains many
// properties to control the behavior of the particles it generates,
// including:
//
// - The birth rate and lifetime of the particle. You can also specify the
// order in which the particles are rendered and the maximum number of
// particles that are created before the emitter turns itself off. - The
// starting values of the particle, including its position, orientation,
// color, and size. You can choose to have these starting values randomized. -
// The changes to apply to the particle over its lifetime. Typically, these
// are specified as a rate of change over time. For example, you might specify
// that a particle rotates at a particular rate, in radians per second. The
// emitter automatically updates the particle data for each frame. In most
// cases, you can also create more sophisticated behaviors using keyframe
// sequences. For example, you might specify a keyframe sequence for a
// particle so that it starts out small, scales up to a larger size, then
// shrinks before dying.
//
// # Choosing Which Node in the Scene Emits Particles
//
//   - [SKEmitterNode.TargetNode]: The target node that renders the emitter’s particles.
//   - [SKEmitterNode.SetTargetNode]
//
// # Controlling When Particles Are Created
//
//   - [SKEmitterNode.AdvanceSimulationTime]: Advances the emitter particle simulation.
//   - [SKEmitterNode.ResetSimulation]: Removes all existing particles and restarts the simulation.
//   - [SKEmitterNode.ParticleBirthRate]: The rate at which new particles are created.
//   - [SKEmitterNode.SetParticleBirthRate]
//   - [SKEmitterNode.NumParticlesToEmit]: The number of particles the emitter should emit before stopping.
//   - [SKEmitterNode.SetNumParticlesToEmit]
//
// # Controlling the Rendering Order of an Emitter’s Particles
//
//   - [SKEmitterNode.ParticleRenderOrder]: The order in which the emitter’s particles are rendered.
//   - [SKEmitterNode.SetParticleRenderOrder]
//
// # Controlling Particle Lifetime
//
//   - [SKEmitterNode.ParticleLifetime]: The average lifetime of a particle, in seconds.
//   - [SKEmitterNode.SetParticleLifetime]
//   - [SKEmitterNode.ParticleLifetimeRange]: The range of allowed random values for a particle’s lifetime.
//   - [SKEmitterNode.SetParticleLifetimeRange]
//
// # Controlling Particle Position
//
//   - [SKEmitterNode.ParticlePosition]: The average starting position for a particle.
//   - [SKEmitterNode.SetParticlePosition]
//   - [SKEmitterNode.ParticlePositionRange]: The range of allowed random values for a particle’s position.
//   - [SKEmitterNode.SetParticlePositionRange]
//   - [SKEmitterNode.ParticleZPosition]: The average starting depth of a particle.
//   - [SKEmitterNode.SetParticleZPosition]
//
// # Controlling Particle Velocity and Acceleration
//
//   - [SKEmitterNode.ParticleSpeed]: The average initial speed of a new particle, in points per second.
//   - [SKEmitterNode.SetParticleSpeed]
//   - [SKEmitterNode.ParticleSpeedRange]: The range of allowed random values for a particle’s initial speed.
//   - [SKEmitterNode.SetParticleSpeedRange]
//   - [SKEmitterNode.EmissionAngle]: The average initial direction of a particle, expressed as an angle in radians.
//   - [SKEmitterNode.SetEmissionAngle]
//   - [SKEmitterNode.EmissionAngleRange]: The range of allowed random values for a particle’s initial direction, expressed as an angle in radians.
//   - [SKEmitterNode.SetEmissionAngleRange]
//   - [SKEmitterNode.XAcceleration]: The acceleration to apply to a particle’s horizontal velocity.
//   - [SKEmitterNode.SetXAcceleration]
//   - [SKEmitterNode.YAcceleration]: The acceleration to apply to a particle’s vertical velocity.
//   - [SKEmitterNode.SetYAcceleration]
//
// # Adjusting a Particle’s Rotation
//
//   - [SKEmitterNode.ParticleRotation]: The average initial rotation of a particle, expressed as an angle in radians.
//   - [SKEmitterNode.SetParticleRotation]
//   - [SKEmitterNode.ParticleRotationRange]: The range of allowed random values for a particle’s initial rotation, expressed as an angle in radians.
//   - [SKEmitterNode.SetParticleRotationRange]
//   - [SKEmitterNode.ParticleRotationSpeed]: The speed at which a particle rotates, expressed in radians per second.
//   - [SKEmitterNode.SetParticleRotationSpeed]
//
// # Scaling Particles by a Factor
//
//   - [SKEmitterNode.ParticleScale]: The average initial scale factor of a particle.
//   - [SKEmitterNode.SetParticleScale]
//   - [SKEmitterNode.ParticleScaleRange]: The range of allowed random values for a particle’s initial scale.
//   - [SKEmitterNode.SetParticleScaleRange]
//   - [SKEmitterNode.ParticleScaleSpeed]: The rate at which a particle’s scale factor changes per second.
//   - [SKEmitterNode.SetParticleScaleSpeed]
//   - [SKEmitterNode.ParticleScaleSequence]: The sequence used to specify the scale factor of a particle over its lifetime.
//   - [SKEmitterNode.SetParticleScaleSequence]
//
// # Changing a Particle’s Source Image and Size
//
//   - [SKEmitterNode.ParticleTexture]: The texture to use to render a particle.
//   - [SKEmitterNode.SetParticleTexture]
//   - [SKEmitterNode.ParticleSize]: The starting size of each particle.
//   - [SKEmitterNode.SetParticleSize]
//
// # Configuring Particle Color
//
//   - [SKEmitterNode.ParticleColorSequence]: The sequence used to specify the color components of a particle over its lifetime.
//   - [SKEmitterNode.SetParticleColorSequence]
//   - [SKEmitterNode.ParticleColor]: The average initial color for a particle.
//   - [SKEmitterNode.SetParticleColor]
//   - [SKEmitterNode.ParticleColorAlphaRange]: The range of allowed random values for the alpha component of a particle’s initial color.
//   - [SKEmitterNode.SetParticleColorAlphaRange]
//   - [SKEmitterNode.ParticleColorBlueRange]: The range of allowed random values for the blue component of a particle’s initial color.
//   - [SKEmitterNode.SetParticleColorBlueRange]
//   - [SKEmitterNode.ParticleColorGreenRange]: The range of allowed random values for the green component of a particle’s initial color.
//   - [SKEmitterNode.SetParticleColorGreenRange]
//   - [SKEmitterNode.ParticleColorRedRange]: The range of allowed random values for the red component of a particle’s initial color.
//   - [SKEmitterNode.SetParticleColorRedRange]
//   - [SKEmitterNode.ParticleColorAlphaSpeed]: The rate at which the alpha component of a particle’s color changes per second.
//   - [SKEmitterNode.SetParticleColorAlphaSpeed]
//   - [SKEmitterNode.ParticleColorBlueSpeed]: The rate at which the blue component of a particle’s color changes per second.
//   - [SKEmitterNode.SetParticleColorBlueSpeed]
//   - [SKEmitterNode.ParticleColorGreenSpeed]: The rate at which the green component of a particle’s color changes per second.
//   - [SKEmitterNode.SetParticleColorGreenSpeed]
//   - [SKEmitterNode.ParticleColorRedSpeed]: The rate at which the red component of a particle’s color changes per second.
//   - [SKEmitterNode.SetParticleColorRedSpeed]
//
// # Controlling How the Texture is Blended with Particle Color
//
//   - [SKEmitterNode.ParticleColorBlendFactorSequence]: The sequence used to specify the color blend factor of a particle over its lifetime.
//   - [SKEmitterNode.SetParticleColorBlendFactorSequence]
//   - [SKEmitterNode.ParticleColorBlendFactor]: The average starting value for the color blend factor.
//   - [SKEmitterNode.SetParticleColorBlendFactor]
//   - [SKEmitterNode.ParticleColorBlendFactorRange]: The range of allowed random values for a particle’s starting color blend factor.
//   - [SKEmitterNode.SetParticleColorBlendFactorRange]
//   - [SKEmitterNode.ParticleColorBlendFactorSpeed]: The rate at which the color blend factor changes per second.
//   - [SKEmitterNode.SetParticleColorBlendFactorSpeed]
//
// # Blending Particles with the Framebuffer
//
//   - [SKEmitterNode.ParticleBlendMode]: The blending mode used to blend particles into the framebuffer.
//   - [SKEmitterNode.SetParticleBlendMode]
//   - [SKEmitterNode.ParticleAlphaSequence]: The sequence used to specify the alpha value of a particle over its lifetime.
//   - [SKEmitterNode.SetParticleAlphaSequence]
//   - [SKEmitterNode.ParticleAlpha]: The average starting alpha value for a particle.
//   - [SKEmitterNode.SetParticleAlpha]
//   - [SKEmitterNode.ParticleAlphaRange]: The range of allowed random values for a particle’s starting alpha value.
//   - [SKEmitterNode.SetParticleAlphaRange]
//   - [SKEmitterNode.ParticleAlphaSpeed]: The rate at which the alpha value of a particle changes per second.
//   - [SKEmitterNode.SetParticleAlphaSpeed]
//
// # Animating Particles
//
//   - [SKEmitterNode.ParticleAction]: An action executed by new particles.
//   - [SKEmitterNode.SetParticleAction]
//
// # Applying Physics Fields to the Particles
//
//   - [SKEmitterNode.FieldBitMask]: A mask that defines which categories of physics fields can exert forces on the particles.
//   - [SKEmitterNode.SetFieldBitMask]
//
// # Taking Full Control of Particle Drawing with a Shader
//
//   - [SKEmitterNode.Shader]: A custom shader used to determine how particles are rendered.
//   - [SKEmitterNode.SetShader]
//   - [SKEmitterNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [SKEmitterNode.SetAttributeValues]
//   - [SKEmitterNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [SKEmitterNode.ValueForAttributeNamed]: Gets the value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode
//
// [SpriteKit]: https://developer.apple.com/documentation/SpriteKit
type SKEmitterNode struct {
	SKNode
}

// SKEmitterNodeFromID constructs a [SKEmitterNode] from an objc.ID.
//
// A source of various particle effects.
func SKEmitterNodeFromID(id objc.ID) SKEmitterNode {
	return SKEmitterNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKEmitterNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKEmitterNode] class.
//
// # Choosing Which Node in the Scene Emits Particles
//
//   - [ISKEmitterNode.TargetNode]: The target node that renders the emitter’s particles.
//   - [ISKEmitterNode.SetTargetNode]
//
// # Controlling When Particles Are Created
//
//   - [ISKEmitterNode.AdvanceSimulationTime]: Advances the emitter particle simulation.
//   - [ISKEmitterNode.ResetSimulation]: Removes all existing particles and restarts the simulation.
//   - [ISKEmitterNode.ParticleBirthRate]: The rate at which new particles are created.
//   - [ISKEmitterNode.SetParticleBirthRate]
//   - [ISKEmitterNode.NumParticlesToEmit]: The number of particles the emitter should emit before stopping.
//   - [ISKEmitterNode.SetNumParticlesToEmit]
//
// # Controlling the Rendering Order of an Emitter’s Particles
//
//   - [ISKEmitterNode.ParticleRenderOrder]: The order in which the emitter’s particles are rendered.
//   - [ISKEmitterNode.SetParticleRenderOrder]
//
// # Controlling Particle Lifetime
//
//   - [ISKEmitterNode.ParticleLifetime]: The average lifetime of a particle, in seconds.
//   - [ISKEmitterNode.SetParticleLifetime]
//   - [ISKEmitterNode.ParticleLifetimeRange]: The range of allowed random values for a particle’s lifetime.
//   - [ISKEmitterNode.SetParticleLifetimeRange]
//
// # Controlling Particle Position
//
//   - [ISKEmitterNode.ParticlePosition]: The average starting position for a particle.
//   - [ISKEmitterNode.SetParticlePosition]
//   - [ISKEmitterNode.ParticlePositionRange]: The range of allowed random values for a particle’s position.
//   - [ISKEmitterNode.SetParticlePositionRange]
//   - [ISKEmitterNode.ParticleZPosition]: The average starting depth of a particle.
//   - [ISKEmitterNode.SetParticleZPosition]
//
// # Controlling Particle Velocity and Acceleration
//
//   - [ISKEmitterNode.ParticleSpeed]: The average initial speed of a new particle, in points per second.
//   - [ISKEmitterNode.SetParticleSpeed]
//   - [ISKEmitterNode.ParticleSpeedRange]: The range of allowed random values for a particle’s initial speed.
//   - [ISKEmitterNode.SetParticleSpeedRange]
//   - [ISKEmitterNode.EmissionAngle]: The average initial direction of a particle, expressed as an angle in radians.
//   - [ISKEmitterNode.SetEmissionAngle]
//   - [ISKEmitterNode.EmissionAngleRange]: The range of allowed random values for a particle’s initial direction, expressed as an angle in radians.
//   - [ISKEmitterNode.SetEmissionAngleRange]
//   - [ISKEmitterNode.XAcceleration]: The acceleration to apply to a particle’s horizontal velocity.
//   - [ISKEmitterNode.SetXAcceleration]
//   - [ISKEmitterNode.YAcceleration]: The acceleration to apply to a particle’s vertical velocity.
//   - [ISKEmitterNode.SetYAcceleration]
//
// # Adjusting a Particle’s Rotation
//
//   - [ISKEmitterNode.ParticleRotation]: The average initial rotation of a particle, expressed as an angle in radians.
//   - [ISKEmitterNode.SetParticleRotation]
//   - [ISKEmitterNode.ParticleRotationRange]: The range of allowed random values for a particle’s initial rotation, expressed as an angle in radians.
//   - [ISKEmitterNode.SetParticleRotationRange]
//   - [ISKEmitterNode.ParticleRotationSpeed]: The speed at which a particle rotates, expressed in radians per second.
//   - [ISKEmitterNode.SetParticleRotationSpeed]
//
// # Scaling Particles by a Factor
//
//   - [ISKEmitterNode.ParticleScale]: The average initial scale factor of a particle.
//   - [ISKEmitterNode.SetParticleScale]
//   - [ISKEmitterNode.ParticleScaleRange]: The range of allowed random values for a particle’s initial scale.
//   - [ISKEmitterNode.SetParticleScaleRange]
//   - [ISKEmitterNode.ParticleScaleSpeed]: The rate at which a particle’s scale factor changes per second.
//   - [ISKEmitterNode.SetParticleScaleSpeed]
//   - [ISKEmitterNode.ParticleScaleSequence]: The sequence used to specify the scale factor of a particle over its lifetime.
//   - [ISKEmitterNode.SetParticleScaleSequence]
//
// # Changing a Particle’s Source Image and Size
//
//   - [ISKEmitterNode.ParticleTexture]: The texture to use to render a particle.
//   - [ISKEmitterNode.SetParticleTexture]
//   - [ISKEmitterNode.ParticleSize]: The starting size of each particle.
//   - [ISKEmitterNode.SetParticleSize]
//
// # Configuring Particle Color
//
//   - [ISKEmitterNode.ParticleColorSequence]: The sequence used to specify the color components of a particle over its lifetime.
//   - [ISKEmitterNode.SetParticleColorSequence]
//   - [ISKEmitterNode.ParticleColor]: The average initial color for a particle.
//   - [ISKEmitterNode.SetParticleColor]
//   - [ISKEmitterNode.ParticleColorAlphaRange]: The range of allowed random values for the alpha component of a particle’s initial color.
//   - [ISKEmitterNode.SetParticleColorAlphaRange]
//   - [ISKEmitterNode.ParticleColorBlueRange]: The range of allowed random values for the blue component of a particle’s initial color.
//   - [ISKEmitterNode.SetParticleColorBlueRange]
//   - [ISKEmitterNode.ParticleColorGreenRange]: The range of allowed random values for the green component of a particle’s initial color.
//   - [ISKEmitterNode.SetParticleColorGreenRange]
//   - [ISKEmitterNode.ParticleColorRedRange]: The range of allowed random values for the red component of a particle’s initial color.
//   - [ISKEmitterNode.SetParticleColorRedRange]
//   - [ISKEmitterNode.ParticleColorAlphaSpeed]: The rate at which the alpha component of a particle’s color changes per second.
//   - [ISKEmitterNode.SetParticleColorAlphaSpeed]
//   - [ISKEmitterNode.ParticleColorBlueSpeed]: The rate at which the blue component of a particle’s color changes per second.
//   - [ISKEmitterNode.SetParticleColorBlueSpeed]
//   - [ISKEmitterNode.ParticleColorGreenSpeed]: The rate at which the green component of a particle’s color changes per second.
//   - [ISKEmitterNode.SetParticleColorGreenSpeed]
//   - [ISKEmitterNode.ParticleColorRedSpeed]: The rate at which the red component of a particle’s color changes per second.
//   - [ISKEmitterNode.SetParticleColorRedSpeed]
//
// # Controlling How the Texture is Blended with Particle Color
//
//   - [ISKEmitterNode.ParticleColorBlendFactorSequence]: The sequence used to specify the color blend factor of a particle over its lifetime.
//   - [ISKEmitterNode.SetParticleColorBlendFactorSequence]
//   - [ISKEmitterNode.ParticleColorBlendFactor]: The average starting value for the color blend factor.
//   - [ISKEmitterNode.SetParticleColorBlendFactor]
//   - [ISKEmitterNode.ParticleColorBlendFactorRange]: The range of allowed random values for a particle’s starting color blend factor.
//   - [ISKEmitterNode.SetParticleColorBlendFactorRange]
//   - [ISKEmitterNode.ParticleColorBlendFactorSpeed]: The rate at which the color blend factor changes per second.
//   - [ISKEmitterNode.SetParticleColorBlendFactorSpeed]
//
// # Blending Particles with the Framebuffer
//
//   - [ISKEmitterNode.ParticleBlendMode]: The blending mode used to blend particles into the framebuffer.
//   - [ISKEmitterNode.SetParticleBlendMode]
//   - [ISKEmitterNode.ParticleAlphaSequence]: The sequence used to specify the alpha value of a particle over its lifetime.
//   - [ISKEmitterNode.SetParticleAlphaSequence]
//   - [ISKEmitterNode.ParticleAlpha]: The average starting alpha value for a particle.
//   - [ISKEmitterNode.SetParticleAlpha]
//   - [ISKEmitterNode.ParticleAlphaRange]: The range of allowed random values for a particle’s starting alpha value.
//   - [ISKEmitterNode.SetParticleAlphaRange]
//   - [ISKEmitterNode.ParticleAlphaSpeed]: The rate at which the alpha value of a particle changes per second.
//   - [ISKEmitterNode.SetParticleAlphaSpeed]
//
// # Animating Particles
//
//   - [ISKEmitterNode.ParticleAction]: An action executed by new particles.
//   - [ISKEmitterNode.SetParticleAction]
//
// # Applying Physics Fields to the Particles
//
//   - [ISKEmitterNode.FieldBitMask]: A mask that defines which categories of physics fields can exert forces on the particles.
//   - [ISKEmitterNode.SetFieldBitMask]
//
// # Taking Full Control of Particle Drawing with a Shader
//
//   - [ISKEmitterNode.Shader]: A custom shader used to determine how particles are rendered.
//   - [ISKEmitterNode.SetShader]
//   - [ISKEmitterNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [ISKEmitterNode.SetAttributeValues]
//   - [ISKEmitterNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [ISKEmitterNode.ValueForAttributeNamed]: Gets the value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode
type ISKEmitterNode interface {
	ISKNode

	// Topic: Choosing Which Node in the Scene Emits Particles

	// The target node that renders the emitter’s particles.
	TargetNode() ISKNode
	SetTargetNode(value ISKNode)

	// Topic: Controlling When Particles Are Created

	// Advances the emitter particle simulation.
	AdvanceSimulationTime(sec foundation.NSTimeInterval)
	// Removes all existing particles and restarts the simulation.
	ResetSimulation()
	// The rate at which new particles are created.
	ParticleBirthRate() float64
	SetParticleBirthRate(value float64)
	// The number of particles the emitter should emit before stopping.
	NumParticlesToEmit() uint
	SetNumParticlesToEmit(value uint)

	// Topic: Controlling the Rendering Order of an Emitter’s Particles

	// The order in which the emitter’s particles are rendered.
	ParticleRenderOrder() SKParticleRenderOrder
	SetParticleRenderOrder(value SKParticleRenderOrder)

	// Topic: Controlling Particle Lifetime

	// The average lifetime of a particle, in seconds.
	ParticleLifetime() float64
	SetParticleLifetime(value float64)
	// The range of allowed random values for a particle’s lifetime.
	ParticleLifetimeRange() float64
	SetParticleLifetimeRange(value float64)

	// Topic: Controlling Particle Position

	// The average starting position for a particle.
	ParticlePosition() corefoundation.CGPoint
	SetParticlePosition(value corefoundation.CGPoint)
	// The range of allowed random values for a particle’s position.
	ParticlePositionRange() corefoundation.CGVector
	SetParticlePositionRange(value corefoundation.CGVector)
	// The average starting depth of a particle.
	ParticleZPosition() float64
	SetParticleZPosition(value float64)

	// Topic: Controlling Particle Velocity and Acceleration

	// The average initial speed of a new particle, in points per second.
	ParticleSpeed() float64
	SetParticleSpeed(value float64)
	// The range of allowed random values for a particle’s initial speed.
	ParticleSpeedRange() float64
	SetParticleSpeedRange(value float64)
	// The average initial direction of a particle, expressed as an angle in radians.
	EmissionAngle() float64
	SetEmissionAngle(value float64)
	// The range of allowed random values for a particle’s initial direction, expressed as an angle in radians.
	EmissionAngleRange() float64
	SetEmissionAngleRange(value float64)
	// The acceleration to apply to a particle’s horizontal velocity.
	XAcceleration() float64
	SetXAcceleration(value float64)
	// The acceleration to apply to a particle’s vertical velocity.
	YAcceleration() float64
	SetYAcceleration(value float64)

	// Topic: Adjusting a Particle’s Rotation

	// The average initial rotation of a particle, expressed as an angle in radians.
	ParticleRotation() float64
	SetParticleRotation(value float64)
	// The range of allowed random values for a particle’s initial rotation, expressed as an angle in radians.
	ParticleRotationRange() float64
	SetParticleRotationRange(value float64)
	// The speed at which a particle rotates, expressed in radians per second.
	ParticleRotationSpeed() float64
	SetParticleRotationSpeed(value float64)

	// Topic: Scaling Particles by a Factor

	// The average initial scale factor of a particle.
	ParticleScale() float64
	SetParticleScale(value float64)
	// The range of allowed random values for a particle’s initial scale.
	ParticleScaleRange() float64
	SetParticleScaleRange(value float64)
	// The rate at which a particle’s scale factor changes per second.
	ParticleScaleSpeed() float64
	SetParticleScaleSpeed(value float64)
	// The sequence used to specify the scale factor of a particle over its lifetime.
	ParticleScaleSequence() ISKKeyframeSequence
	SetParticleScaleSequence(value ISKKeyframeSequence)

	// Topic: Changing a Particle’s Source Image and Size

	// The texture to use to render a particle.
	ParticleTexture() ISKTexture
	SetParticleTexture(value ISKTexture)
	// The starting size of each particle.
	ParticleSize() corefoundation.CGSize
	SetParticleSize(value corefoundation.CGSize)

	// Topic: Configuring Particle Color

	// The sequence used to specify the color components of a particle over its lifetime.
	ParticleColorSequence() ISKKeyframeSequence
	SetParticleColorSequence(value ISKKeyframeSequence)
	// The average initial color for a particle.
	ParticleColor() appkit.NSColor
	SetParticleColor(value appkit.NSColor)
	// The range of allowed random values for the alpha component of a particle’s initial color.
	ParticleColorAlphaRange() float64
	SetParticleColorAlphaRange(value float64)
	// The range of allowed random values for the blue component of a particle’s initial color.
	ParticleColorBlueRange() float64
	SetParticleColorBlueRange(value float64)
	// The range of allowed random values for the green component of a particle’s initial color.
	ParticleColorGreenRange() float64
	SetParticleColorGreenRange(value float64)
	// The range of allowed random values for the red component of a particle’s initial color.
	ParticleColorRedRange() float64
	SetParticleColorRedRange(value float64)
	// The rate at which the alpha component of a particle’s color changes per second.
	ParticleColorAlphaSpeed() float64
	SetParticleColorAlphaSpeed(value float64)
	// The rate at which the blue component of a particle’s color changes per second.
	ParticleColorBlueSpeed() float64
	SetParticleColorBlueSpeed(value float64)
	// The rate at which the green component of a particle’s color changes per second.
	ParticleColorGreenSpeed() float64
	SetParticleColorGreenSpeed(value float64)
	// The rate at which the red component of a particle’s color changes per second.
	ParticleColorRedSpeed() float64
	SetParticleColorRedSpeed(value float64)

	// Topic: Controlling How the Texture is Blended with Particle Color

	// The sequence used to specify the color blend factor of a particle over its lifetime.
	ParticleColorBlendFactorSequence() ISKKeyframeSequence
	SetParticleColorBlendFactorSequence(value ISKKeyframeSequence)
	// The average starting value for the color blend factor.
	ParticleColorBlendFactor() float64
	SetParticleColorBlendFactor(value float64)
	// The range of allowed random values for a particle’s starting color blend factor.
	ParticleColorBlendFactorRange() float64
	SetParticleColorBlendFactorRange(value float64)
	// The rate at which the color blend factor changes per second.
	ParticleColorBlendFactorSpeed() float64
	SetParticleColorBlendFactorSpeed(value float64)

	// Topic: Blending Particles with the Framebuffer

	// The blending mode used to blend particles into the framebuffer.
	ParticleBlendMode() SKBlendMode
	SetParticleBlendMode(value SKBlendMode)
	// The sequence used to specify the alpha value of a particle over its lifetime.
	ParticleAlphaSequence() ISKKeyframeSequence
	SetParticleAlphaSequence(value ISKKeyframeSequence)
	// The average starting alpha value for a particle.
	ParticleAlpha() float64
	SetParticleAlpha(value float64)
	// The range of allowed random values for a particle’s starting alpha value.
	ParticleAlphaRange() float64
	SetParticleAlphaRange(value float64)
	// The rate at which the alpha value of a particle changes per second.
	ParticleAlphaSpeed() float64
	SetParticleAlphaSpeed(value float64)

	// Topic: Animating Particles

	// An action executed by new particles.
	ParticleAction() ISKAction
	SetParticleAction(value ISKAction)

	// Topic: Applying Physics Fields to the Particles

	// A mask that defines which categories of physics fields can exert forces on the particles.
	FieldBitMask() uint32
	SetFieldBitMask(value uint32)

	// Topic: Taking Full Control of Particle Drawing with a Shader

	// A custom shader used to determine how particles are rendered.
	Shader() ISKShader
	SetShader(value ISKShader)
	// The values of each attribute associated with the node’s attached shader.
	AttributeValues() foundation.INSDictionary
	SetAttributeValues(value foundation.INSDictionary)
	// Sets an attribute value for an attached shader.
	SetValueForAttributeNamed(value ISKAttributeValue, key string)
	// Gets the value of a shader attribute.
	ValueForAttributeNamed(key string) ISKAttributeValue
}

// Init initializes the instance.
func (e SKEmitterNode) Init() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e SKEmitterNode) Autorelease() SKEmitterNode {
	rv := objc.Send[SKEmitterNode](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKEmitterNode creates a new SKEmitterNode instance.
func NewSKEmitterNode() SKEmitterNode {
	class := getSKEmitterNodeClass()
	rv := objc.Send[SKEmitterNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewEmitterNodeWithCoder(aDecoder foundation.INSCoder) SKEmitterNode {
	instance := getSKEmitterNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKEmitterNodeFromID(rv)
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewEmitterNodeWithFileNamed(filename string) SKEmitterNode {
	rv := objc.Send[objc.ID](objc.ID(getSKEmitterNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKEmitterNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewEmitterNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKEmitterNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKEmitterNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKEmitterNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKEmitterNode{}, objc.ErrInitFailed
	}
	return SKEmitterNodeFromID(rv), nil
}

// Advances the emitter particle simulation.
//
// sec: The number of seconds to simulate.
//
// # Discussion
//
// Once added to a scene, an emitter node automatically creates new particles
// in new animation frames. This method allows you to artificially advance a
// running emitter’s simulation, causing it to generate new particles and
// advance any existing particles. The most common use for this method is to
// prepopulate an emitter node with particles after it is first added to a
// scene.
//
// If an emitter is paused, either directly or by one of its parent nodes
// being paused, [SKEmitterNode.AdvanceSimulationTime] has no effect. Listing
// 1 shows how to advance an emitter object while it is paused, by temporarily
// toggling its [SKNode.Paused] property while advancing the simulation.
//
// Listing 1. Temporarily unpausing an emitter to advance simulation time
//
// Note that the [SKNode.Paused] property is inherited from a node’s
// parents, so even if the emitter hasn’t been explicitly paused but, for
// example, the scene has, this code will still work.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/advanceSimulationTime(_:)
func (e SKEmitterNode) AdvanceSimulationTime(sec foundation.NSTimeInterval) {
	objc.Send[objc.ID](e.ID, objc.Sel("advanceSimulationTime:"), sec)
}

// Removes all existing particles and restarts the simulation.
//
// # Discussion
//
// Resetting the simulation clears the internal state of the simulation.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/resetSimulation()
func (e SKEmitterNode) ResetSimulation() {
	objc.Send[objc.ID](e.ID, objc.Sel("resetSimulation"))
}

// Sets an attribute value for an attached shader.
//
// value: An attribute value object containing the scalar or vector value to set in
// the attached shader.
//
// key: The attribute name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/setValue(_:forAttribute:)
func (e SKEmitterNode) SetValueForAttributeNamed(value ISKAttributeValue, key string) {
	objc.Send[objc.ID](e.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}

// Gets the value of a shader attribute.
//
// key: The attribute name.
//
// # Return Value
//
// An attribute value object containing the scalar or vector value or `nil` if
// no such attribute exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/value(forAttributeNamed:)
func (e SKEmitterNode) ValueForAttributeNamed(key string) ISKAttributeValue {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return SKAttributeValueFromID(rv)
}

// The target node that renders the emitter’s particles.
//
// # Discussion
//
// The default value is `nil`, which means that particles are treated as if
// they were children of the emitter node. In future frames of animation, the
// particle positions are affected by the emitter node’s properties. If you
// specify a different target node, the initial properties of new particles
// are calculated based on the emitter node’s properties, but in future
// frames of animation the particles are treated as if they were children of
// the target node.
//
// For example, assume you have an emitter node as a child of the scene node
// and the node is being rotated by changing its [SKNode.ZRotation] property.
// The behavior of the emitter node changes based on the value of the target
// node:
//
// - If the [SKEmitterNode.TargetNode] property is `nil`, then the positions
// of both previously generated and new particles are rotated. - If the
// [SKEmitterNode.TargetNode] property points to the scene node, then new
// particles are adjusted when the emitter node rotates, but previously
// generated particles are not.
//
// By spawning the particles inside the scene node, they have behavior
// independent of the emitter’s properties.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/targetNode
func (e SKEmitterNode) TargetNode() ISKNode {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("targetNode"))
	return SKNodeFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetTargetNode(value ISKNode) {
	objc.Send[struct{}](e.ID, objc.Sel("setTargetNode:"), value)
}

// The rate at which new particles are created.
//
// # Discussion
//
// The number of particles generated by the emitter every second. The default
// value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleBirthRate
func (e SKEmitterNode) ParticleBirthRate() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleBirthRate"))
	return rv
}
func (e SKEmitterNode) SetParticleBirthRate(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleBirthRate:"), value)
}

// The number of particles the emitter should emit before stopping.
//
// # Discussion
//
// The default value is `0`, which indicates that emitter creates an endless
// stream of particles. If a non-zero value is provided, then the emitter
// stops generating particles after it has created the specified number of
// particles.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/numParticlesToEmit
func (e SKEmitterNode) NumParticlesToEmit() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("numParticlesToEmit"))
	return rv
}
func (e SKEmitterNode) SetNumParticlesToEmit(value uint) {
	objc.Send[struct{}](e.ID, objc.Sel("setNumParticlesToEmit:"), value)
}

// The order in which the emitter’s particles are rendered.
//
// # Discussion
//
// The default value is [SKParticleRenderOrder.oldestLast].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleRenderOrder
//
// [SKParticleRenderOrder.oldestLast]: https://developer.apple.com/documentation/SpriteKit/SKParticleRenderOrder/oldestLast
func (e SKEmitterNode) ParticleRenderOrder() SKParticleRenderOrder {
	rv := objc.Send[SKParticleRenderOrder](e.ID, objc.Sel("particleRenderOrder"))
	return SKParticleRenderOrder(rv)
}
func (e SKEmitterNode) SetParticleRenderOrder(value SKParticleRenderOrder) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleRenderOrder:"), value)
}

// The average lifetime of a particle, in seconds.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleLifetime
func (e SKEmitterNode) ParticleLifetime() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleLifetime"))
	return rv
}
func (e SKEmitterNode) SetParticleLifetime(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleLifetime:"), value)
}

// The range of allowed random values for a particle’s lifetime.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the lifetime of each particle is
// randomly determined and may vary by plus or minus half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleLifetimeRange
func (e SKEmitterNode) ParticleLifetimeRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleLifetimeRange"))
	return rv
}
func (e SKEmitterNode) SetParticleLifetimeRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleLifetimeRange:"), value)
}

// The average starting position for a particle.
//
// # Discussion
//
// The default value is `(0.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particlePosition
func (e SKEmitterNode) ParticlePosition() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](e.ID, objc.Sel("particlePosition"))
	return corefoundation.CGPoint(rv)
}
func (e SKEmitterNode) SetParticlePosition(value corefoundation.CGPoint) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticlePosition:"), value)
}

// The range of allowed random values for a particle’s position.
//
// # Discussion
//
// The default value is `(0.0,0.0)`. If a component is non-zero, the same
// component of a particle’s position is randomly determined and may vary by
// plus or minus half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particlePositionRange
func (e SKEmitterNode) ParticlePositionRange() corefoundation.CGVector {
	rv := objc.Send[corefoundation.CGVector](e.ID, objc.Sel("particlePositionRange"))
	return corefoundation.CGVector(rv)
}
func (e SKEmitterNode) SetParticlePositionRange(value corefoundation.CGVector) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticlePositionRange:"), value)
}

// The average starting depth of a particle.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleZPosition
func (e SKEmitterNode) ParticleZPosition() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleZPosition"))
	return rv
}
func (e SKEmitterNode) SetParticleZPosition(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleZPosition:"), value)
}

// The average initial speed of a new particle, in points per second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleSpeed
func (e SKEmitterNode) ParticleSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleSpeed:"), value)
}

// The range of allowed random values for a particle’s initial speed.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the speed of each particle is
// randomly determined and may vary by plus or minus half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleSpeedRange
func (e SKEmitterNode) ParticleSpeedRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleSpeedRange"))
	return rv
}
func (e SKEmitterNode) SetParticleSpeedRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleSpeedRange:"), value)
}

// The average initial direction of a particle, expressed as an angle in
// radians.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/emissionAngle
func (e SKEmitterNode) EmissionAngle() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionAngle"))
	return rv
}
func (e SKEmitterNode) SetEmissionAngle(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionAngle:"), value)
}

// The range of allowed random values for a particle’s initial direction,
// expressed as an angle in radians.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the emission angle of each
// particle is randomly determined and may vary by plus or minus half of the
// range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/emissionAngleRange
func (e SKEmitterNode) EmissionAngleRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("emissionAngleRange"))
	return rv
}
func (e SKEmitterNode) SetEmissionAngleRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setEmissionAngleRange:"), value)
}

// The acceleration to apply to a particle’s horizontal velocity.
//
// # Discussion
//
// This property is useful for simulating wind, gravity and other effects. It
// is uniformly applied to all particles generated by the emitter.
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/xAcceleration
func (e SKEmitterNode) XAcceleration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("xAcceleration"))
	return rv
}
func (e SKEmitterNode) SetXAcceleration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setXAcceleration:"), value)
}

// The acceleration to apply to a particle’s vertical velocity.
//
// # Discussion
//
// This property is useful for simulating wind, gravity and other effects. It
// is uniformly applied to all particles generated by the emitter.
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/yAcceleration
func (e SKEmitterNode) YAcceleration() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("yAcceleration"))
	return rv
}
func (e SKEmitterNode) SetYAcceleration(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setYAcceleration:"), value)
}

// The average initial rotation of a particle, expressed as an angle in
// radians.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleRotation
func (e SKEmitterNode) ParticleRotation() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleRotation"))
	return rv
}
func (e SKEmitterNode) SetParticleRotation(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleRotation:"), value)
}

// The range of allowed random values for a particle’s initial rotation,
// expressed as an angle in radians.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the initial rotation of each
// particle is randomly determined and may vary by plus or minus half of the
// range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleRotationRange
func (e SKEmitterNode) ParticleRotationRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleRotationRange"))
	return rv
}
func (e SKEmitterNode) SetParticleRotationRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleRotationRange:"), value)
}

// The speed at which a particle rotates, expressed in radians per second.
//
// # Discussion
//
// This is uniform for all particles generated by the emitter. The default
// value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleRotationSpeed
func (e SKEmitterNode) ParticleRotationSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleRotationSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleRotationSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleRotationSpeed:"), value)
}

// The average initial scale factor of a particle.
//
// # Discussion
//
// The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleScale
func (e SKEmitterNode) ParticleScale() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleScale"))
	return rv
}
func (e SKEmitterNode) SetParticleScale(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleScale:"), value)
}

// The range of allowed random values for a particle’s initial scale.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the initial scale of each particle
// is randomly determined and may vary by plus or minus half of the range
// value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleScaleRange
func (e SKEmitterNode) ParticleScaleRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleScaleRange"))
	return rv
}
func (e SKEmitterNode) SetParticleScaleRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleScaleRange:"), value)
}

// The rate at which a particle’s scale factor changes per second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleScaleSpeed
func (e SKEmitterNode) ParticleScaleSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleScaleSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleScaleSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleScaleSpeed:"), value)
}

// The sequence used to specify the scale factor of a particle over its
// lifetime.
//
// # Discussion
//
// The default value is `nil`. If a non-`nil` value is specified, then the
// [SKEmitterNode.ParticleScale], [SKEmitterNode.ParticleScaleRange], and
// [SKEmitterNode.ParticleScaleSpeed] properties are ignored. Instead, the
// sequence is used to specify the scale factor.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleScaleSequence
func (e SKEmitterNode) ParticleScaleSequence() ISKKeyframeSequence {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleScaleSequence"))
	return SKKeyframeSequenceFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleScaleSequence(value ISKKeyframeSequence) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleScaleSequence:"), value)
}

// The texture to use to render a particle.
//
// # Discussion
//
// A particle is rendered as if it were a [SKSpriteNode] object. The default
// value is `nil`, which means a single-color rectangle is used to draw the
// particle. If a non-`nil` value is specified, then the texture is colorized
// and used to draw particles.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleTexture
func (e SKEmitterNode) ParticleTexture() ISKTexture {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleTexture"))
	return SKTextureFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleTexture(value ISKTexture) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleTexture:"), value)
}

// The starting size of each particle.
//
// # Discussion
//
// The default value is [CGSizeZero]. If set to the default, the size of the
// texture stored in the [SKEmitterNode.ParticleTexture] property is used to
// determine the size of a particle. If a texture has not been assigned, you
// must set this property to a non-empty size.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleSize
//
// [CGSizeZero]: https://developer.apple.com/documentation/CoreGraphics/CGSizeZero
func (e SKEmitterNode) ParticleSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e.ID, objc.Sel("particleSize"))
	return corefoundation.CGSize(rv)
}
func (e SKEmitterNode) SetParticleSize(value corefoundation.CGSize) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleSize:"), value)
}

// The sequence used to specify the color components of a particle over its
// lifetime.
//
// # Discussion
//
// The default value is `nil`. If a non-`nil` value is specified, then the
// [SKEmitterNode.ParticleColor], [SKEmitterNode.ParticleColorAlphaRange],
// [SKEmitterNode.ParticleColorRedRange],
// [SKEmitterNode.ParticleColorGreenRange],
// [SKEmitterNode.ParticleColorBlueRange],
// [SKEmitterNode.ParticleColorAlphaSpeed],
// [SKEmitterNode.ParticleColorRedSpeed],
// [SKEmitterNode.ParticleColorGreenSpeed], and
// [SKEmitterNode.ParticleColorBlueSpeed] properties are ignored. Instead, the
// sequence is used to specify the particle color.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorSequence
func (e SKEmitterNode) ParticleColorSequence() ISKKeyframeSequence {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleColorSequence"))
	return SKKeyframeSequenceFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleColorSequence(value ISKKeyframeSequence) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorSequence:"), value)
}

// The average initial color for a particle.
//
// # Discussion
//
// The default value is `[SKColor clearColor]`.
//
// A particle’s color is blended with the texture using its blend color
// factor. See [SKEmitterNode].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColor
func (e SKEmitterNode) ParticleColor() appkit.NSColor {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleColor(value appkit.NSColor) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColor:"), value)
}

// The range of allowed random values for the alpha component of a
// particle’s initial color.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the starting alpha component of a
// particle’s color is randomly determined and may vary by plus or minus
// half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorAlphaRange
func (e SKEmitterNode) ParticleColorAlphaRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorAlphaRange"))
	return rv
}
func (e SKEmitterNode) SetParticleColorAlphaRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorAlphaRange:"), value)
}

// The range of allowed random values for the blue component of a particle’s
// initial color.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the starting blue component of a
// particle’s color is randomly determined and may vary by plus or minus
// half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlueRange
func (e SKEmitterNode) ParticleColorBlueRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorBlueRange"))
	return rv
}
func (e SKEmitterNode) SetParticleColorBlueRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlueRange:"), value)
}

// The range of allowed random values for the green component of a
// particle’s initial color.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the starting green component of a
// particle’s color is randomly determined and may vary by plus or minus
// half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorGreenRange
func (e SKEmitterNode) ParticleColorGreenRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorGreenRange"))
	return rv
}
func (e SKEmitterNode) SetParticleColorGreenRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorGreenRange:"), value)
}

// The range of allowed random values for the red component of a particle’s
// initial color.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the starting red component of a
// particle’s color is randomly determined and may vary by plus or minus
// half of the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorRedRange
func (e SKEmitterNode) ParticleColorRedRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorRedRange"))
	return rv
}
func (e SKEmitterNode) SetParticleColorRedRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorRedRange:"), value)
}

// The rate at which the alpha component of a particle’s color changes per
// second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorAlphaSpeed
func (e SKEmitterNode) ParticleColorAlphaSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorAlphaSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleColorAlphaSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorAlphaSpeed:"), value)
}

// The rate at which the blue component of a particle’s color changes per
// second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlueSpeed
func (e SKEmitterNode) ParticleColorBlueSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorBlueSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleColorBlueSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlueSpeed:"), value)
}

// The rate at which the green component of a particle’s color changes per
// second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorGreenSpeed
func (e SKEmitterNode) ParticleColorGreenSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorGreenSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleColorGreenSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorGreenSpeed:"), value)
}

// The rate at which the red component of a particle’s color changes per
// second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorRedSpeed
func (e SKEmitterNode) ParticleColorRedSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorRedSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleColorRedSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorRedSpeed:"), value)
}

// The sequence used to specify the color blend factor of a particle over its
// lifetime.
//
// # Discussion
//
// The default value is `nil`. If a non-`nil` value is specified, then the
// [SKEmitterNode.ParticleColorBlendFactor],
// [SKEmitterNode.ParticleColorBlendFactorRange], and
// [SKEmitterNode.ParticleColorBlendFactorSpeed] properties are ignored.
// Instead, the sequence is used to specify the color blend factor.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlendFactorSequence
func (e SKEmitterNode) ParticleColorBlendFactorSequence() ISKKeyframeSequence {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleColorBlendFactorSequence"))
	return SKKeyframeSequenceFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleColorBlendFactorSequence(value ISKKeyframeSequence) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlendFactorSequence:"), value)
}

// The average starting value for the color blend factor.
//
// # Discussion
//
// The default value is `0.0`, which means that the texture is used as is,
// ignoring the particle’s color. Otherwise, the texture is blended with the
// color. The maximum value is `1.0`, which means the particle renders with a
// full-tint color.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlendFactor
func (e SKEmitterNode) ParticleColorBlendFactor() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorBlendFactor"))
	return rv
}
func (e SKEmitterNode) SetParticleColorBlendFactor(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlendFactor:"), value)
}

// The range of allowed random values for a particle’s starting color blend
// factor.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the initial color blend factor of
// each particle is randomly determined and may vary by plus or minus half of
// the range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlendFactorRange
func (e SKEmitterNode) ParticleColorBlendFactorRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorBlendFactorRange"))
	return rv
}
func (e SKEmitterNode) SetParticleColorBlendFactorRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlendFactorRange:"), value)
}

// The rate at which the color blend factor changes per second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleColorBlendFactorSpeed
func (e SKEmitterNode) ParticleColorBlendFactorSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleColorBlendFactorSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleColorBlendFactorSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleColorBlendFactorSpeed:"), value)
}

// The blending mode used to blend particles into the framebuffer.
//
// # Discussion
//
// The default value is [SKBlendMode.alpha].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleBlendMode
//
// [SKBlendMode.alpha]: https://developer.apple.com/documentation/SpriteKit/SKBlendMode/alpha
func (e SKEmitterNode) ParticleBlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](e.ID, objc.Sel("particleBlendMode"))
	return SKBlendMode(rv)
}
func (e SKEmitterNode) SetParticleBlendMode(value SKBlendMode) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleBlendMode:"), value)
}

// The sequence used to specify the alpha value of a particle over its
// lifetime.
//
// # Discussion
//
// The default value is `nil`. If a non-`nil` value is specified, then the
// [SKEmitterNode.ParticleAlpha], [SKEmitterNode.ParticleAlphaRange], and
// [SKEmitterNode.ParticleAlphaSpeed] properties are ignored. Instead, the
// sequence is used to specify the alpha value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleAlphaSequence
func (e SKEmitterNode) ParticleAlphaSequence() ISKKeyframeSequence {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleAlphaSequence"))
	return SKKeyframeSequenceFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleAlphaSequence(value ISKKeyframeSequence) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleAlphaSequence:"), value)
}

// The average starting alpha value for a particle.
//
// # Discussion
//
// The particle alpha value is equivalent to the [SKNode.Alpha] property of
// the [SKNode] class. The alpha component of the color that results from the
// texture and color blending state is multiplied by a particle’s alpha
// value. The resulting particle color is then blended with the parent’s
// framebuffer. The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleAlpha
func (e SKEmitterNode) ParticleAlpha() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleAlpha"))
	return rv
}
func (e SKEmitterNode) SetParticleAlpha(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleAlpha:"), value)
}

// The range of allowed random values for a particle’s starting alpha value.
//
// # Discussion
//
// The default value is `0.0`. If non-zero, the initial alpha value of each
// particle is randomly determined and may vary by plus or minus half of the
// range value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleAlphaRange
func (e SKEmitterNode) ParticleAlphaRange() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleAlphaRange"))
	return rv
}
func (e SKEmitterNode) SetParticleAlphaRange(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleAlphaRange:"), value)
}

// The rate at which the alpha value of a particle changes per second.
//
// # Discussion
//
// The default value is `0.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleAlphaSpeed
func (e SKEmitterNode) ParticleAlphaSpeed() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("particleAlphaSpeed"))
	return rv
}
func (e SKEmitterNode) SetParticleAlphaSpeed(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleAlphaSpeed:"), value)
}

// An action executed by new particles.
//
// # Discussion
//
// Although you do not have direct access to the particles created by
// SpriteKit, you can specify an action that all particles execute. Whenever a
// new particle is created, the emitter tells the particle to run that action.
// You can use actions to create very sophisticated behaviors.
//
// For the purpose of using actions on particles, you can treat the particle
// as if it were a normal node. This means you can perform other interesting
// tricks, such as animating the particle’s textures.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/particleAction
func (e SKEmitterNode) ParticleAction() ISKAction {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("particleAction"))
	return SKActionFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetParticleAction(value ISKAction) {
	objc.Send[struct{}](e.ID, objc.Sel("setParticleAction:"), value)
}

// A mask that defines which categories of physics fields can exert forces on
// the particles.
//
// # Discussion
//
// When a particle is inside the region of a [SKFieldNode] object, that field
// node’s [SKFieldNode.CategoryBitMask] property is compared to the
// emitter’s [SKEmitterNode.FieldBitMask] property by performing a logical
// AND operation. If the result is a non-zero value, then the field node’s
// effect is applied to the particle as if it had a physics body. The physics
// body is assumed to have a [SKPhysicsBody.Mass] of `1.0` and a
// [SKPhysicsBody.Charge] of `1.0`
//
// The default value is `0x00000000` (all bits cleared).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/fieldBitMask
func (e SKEmitterNode) FieldBitMask() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("fieldBitMask"))
	return rv
}
func (e SKEmitterNode) SetFieldBitMask(value uint32) {
	objc.Send[struct{}](e.ID, objc.Sel("setFieldBitMask:"), value)
}

// A custom shader used to determine how particles are rendered.
//
// # Discussion
//
// The default value is `nil`. If a shader is specified, then the shader is
// used to determine the output colors for any of the emitter’s particles.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/shader
func (e SKEmitterNode) Shader() ISKShader {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shader"))
	return SKShaderFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetShader(value ISKShader) {
	objc.Send[struct{}](e.ID, objc.Sel("setShader:"), value)
}

// The values of each attribute associated with the node’s attached shader.
//
// # Discussion
//
// All nodes have their own copy of an attribute value and therefore the
// attribute values can be different per-node across the same [SKShader]. If
// instead you need all nodes to share the same value, use [SKUniform].
// Uniforms can change values every frame, but uniforms cannot vary per-node
// like attributes can.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKEmitterNode/attributeValues
func (e SKEmitterNode) AttributeValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("attributeValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (e SKEmitterNode) SetAttributeValues(value foundation.INSDictionary) {
	objc.Send[struct{}](e.ID, objc.Sel("setAttributeValues:"), value)
}
