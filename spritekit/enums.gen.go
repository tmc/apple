// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/SpriteKit/SKActionTimingMode
type SKActionTimingMode int

const (
	// SKActionTimingEaseIn: Specifies ease-in pacing.
	SKActionTimingEaseIn SKActionTimingMode = 1
	// SKActionTimingEaseInEaseOut: Specifies ease-in ease-out pacing.
	SKActionTimingEaseInEaseOut SKActionTimingMode = 3
	// SKActionTimingEaseOut: Specifies ease-out pacing.
	SKActionTimingEaseOut SKActionTimingMode = 2
	// SKActionTimingLinear: Specifies linear pacing.
	SKActionTimingLinear SKActionTimingMode = 0
)

func (e SKActionTimingMode) String() string {
	switch e {
	case SKActionTimingEaseIn:
		return "SKActionTimingEaseIn"
	case SKActionTimingEaseInEaseOut:
		return "SKActionTimingEaseInEaseOut"
	case SKActionTimingEaseOut:
		return "SKActionTimingEaseOut"
	case SKActionTimingLinear:
		return "SKActionTimingLinear"
	default:
		return fmt.Sprintf("SKActionTimingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKAttributeType
type SKAttributeType int

const (
	SKAttributeTypeFloat            SKAttributeType = 1
	SKAttributeTypeHalfFloat        SKAttributeType = 5
	SKAttributeTypeNone             SKAttributeType = 0
	SKAttributeTypeVectorFloat2     SKAttributeType = 2
	SKAttributeTypeVectorFloat3     SKAttributeType = 3
	SKAttributeTypeVectorFloat4     SKAttributeType = 4
	SKAttributeTypeVectorHalfFloat2 SKAttributeType = 6
	SKAttributeTypeVectorHalfFloat3 SKAttributeType = 7
	SKAttributeTypeVectorHalfFloat4 SKAttributeType = 8
)

func (e SKAttributeType) String() string {
	switch e {
	case SKAttributeTypeFloat:
		return "SKAttributeTypeFloat"
	case SKAttributeTypeHalfFloat:
		return "SKAttributeTypeHalfFloat"
	case SKAttributeTypeNone:
		return "SKAttributeTypeNone"
	case SKAttributeTypeVectorFloat2:
		return "SKAttributeTypeVectorFloat2"
	case SKAttributeTypeVectorFloat3:
		return "SKAttributeTypeVectorFloat3"
	case SKAttributeTypeVectorFloat4:
		return "SKAttributeTypeVectorFloat4"
	case SKAttributeTypeVectorHalfFloat2:
		return "SKAttributeTypeVectorHalfFloat2"
	case SKAttributeTypeVectorHalfFloat3:
		return "SKAttributeTypeVectorHalfFloat3"
	case SKAttributeTypeVectorHalfFloat4:
		return "SKAttributeTypeVectorHalfFloat4"
	default:
		return fmt.Sprintf("SKAttributeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKBlendMode
type SKBlendMode int

const (
	// SKBlendModeAdd: The source and destination colors are added together.
	SKBlendModeAdd SKBlendMode = 1
	// SKBlendModeAlpha: The source and destination colors are blended by multiplying the source alpha value.
	SKBlendModeAlpha SKBlendMode = 0
	// SKBlendModeMultiply: The source color is multiplied by the destination color.
	SKBlendModeMultiply      SKBlendMode = 3
	SKBlendModeMultiplyAlpha SKBlendMode = 7
	// SKBlendModeMultiplyX2: The source color is multiplied by the destination color and then doubled.
	SKBlendModeMultiplyX2 SKBlendMode = 4
	// SKBlendModeReplace: The source color replaces the destination color.
	SKBlendModeReplace SKBlendMode = 6
	// SKBlendModeScreen: The source color is added to the destination color times the inverted source color.
	SKBlendModeScreen SKBlendMode = 5
	// SKBlendModeSubtract: The source color is subtracted from the destination color.
	SKBlendModeSubtract SKBlendMode = 2
)

func (e SKBlendMode) String() string {
	switch e {
	case SKBlendModeAdd:
		return "SKBlendModeAdd"
	case SKBlendModeAlpha:
		return "SKBlendModeAlpha"
	case SKBlendModeMultiply:
		return "SKBlendModeMultiply"
	case SKBlendModeMultiplyAlpha:
		return "SKBlendModeMultiplyAlpha"
	case SKBlendModeMultiplyX2:
		return "SKBlendModeMultiplyX2"
	case SKBlendModeReplace:
		return "SKBlendModeReplace"
	case SKBlendModeScreen:
		return "SKBlendModeScreen"
	case SKBlendModeSubtract:
		return "SKBlendModeSubtract"
	default:
		return fmt.Sprintf("SKBlendMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKInterpolationMode
type SKInterpolationMode int

const (
	// SKInterpolationModeLinear: Values between two keyframes are interpolated linearly.
	SKInterpolationModeLinear SKInterpolationMode = 1
	// SKInterpolationModeSpline: Values between two keyframes using a spline curve.
	SKInterpolationModeSpline SKInterpolationMode = 2
	// SKInterpolationModeStep: Values between two keyframes are not interpolated.
	SKInterpolationModeStep SKInterpolationMode = 3
)

func (e SKInterpolationMode) String() string {
	switch e {
	case SKInterpolationModeLinear:
		return "SKInterpolationModeLinear"
	case SKInterpolationModeSpline:
		return "SKInterpolationModeSpline"
	case SKInterpolationModeStep:
		return "SKInterpolationModeStep"
	default:
		return fmt.Sprintf("SKInterpolationMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKLabelHorizontalAlignmentMode
type SKLabelHorizontalAlignmentMode int

const (
	// SKLabelHorizontalAlignmentModeCenter: Centers the text horizontally on the node’s origin.
	SKLabelHorizontalAlignmentModeCenter SKLabelHorizontalAlignmentMode = 0
	// SKLabelHorizontalAlignmentModeLeft: Positions the text so that the left side of the text is on the node’s origin.
	SKLabelHorizontalAlignmentModeLeft SKLabelHorizontalAlignmentMode = 1
	// SKLabelHorizontalAlignmentModeRight: Positions the text so that the right side of the text is on the node’s origin.
	SKLabelHorizontalAlignmentModeRight SKLabelHorizontalAlignmentMode = 2
)

func (e SKLabelHorizontalAlignmentMode) String() string {
	switch e {
	case SKLabelHorizontalAlignmentModeCenter:
		return "SKLabelHorizontalAlignmentModeCenter"
	case SKLabelHorizontalAlignmentModeLeft:
		return "SKLabelHorizontalAlignmentModeLeft"
	case SKLabelHorizontalAlignmentModeRight:
		return "SKLabelHorizontalAlignmentModeRight"
	default:
		return fmt.Sprintf("SKLabelHorizontalAlignmentMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKLabelVerticalAlignmentMode
type SKLabelVerticalAlignmentMode int

const (
	// SKLabelVerticalAlignmentModeBaseline: Positions the text so that the font’s baseline lies on the node’s origin.
	SKLabelVerticalAlignmentModeBaseline SKLabelVerticalAlignmentMode = 0
	// SKLabelVerticalAlignmentModeBottom: Positions the text so that the bottom of the text is on the node’s origin.
	SKLabelVerticalAlignmentModeBottom SKLabelVerticalAlignmentMode = 3
	// SKLabelVerticalAlignmentModeCenter: Centers the text vertically on the node’s origin.
	SKLabelVerticalAlignmentModeCenter SKLabelVerticalAlignmentMode = 1
	// SKLabelVerticalAlignmentModeTop: Positions the text so that the top of the text is on the node’s origin.
	SKLabelVerticalAlignmentModeTop SKLabelVerticalAlignmentMode = 2
)

func (e SKLabelVerticalAlignmentMode) String() string {
	switch e {
	case SKLabelVerticalAlignmentModeBaseline:
		return "SKLabelVerticalAlignmentModeBaseline"
	case SKLabelVerticalAlignmentModeBottom:
		return "SKLabelVerticalAlignmentModeBottom"
	case SKLabelVerticalAlignmentModeCenter:
		return "SKLabelVerticalAlignmentModeCenter"
	case SKLabelVerticalAlignmentModeTop:
		return "SKLabelVerticalAlignmentModeTop"
	default:
		return fmt.Sprintf("SKLabelVerticalAlignmentMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNodeFocusBehavior
type SKNodeFocusBehavior int

const (
	// SKNodeFocusBehaviorFocusable: Node is focusable and prevents nodes that it visually obscures from becoming focusable.
	SKNodeFocusBehaviorFocusable SKNodeFocusBehavior = 2
	// SKNodeFocusBehaviorNone: Node is not focusable.
	SKNodeFocusBehaviorNone SKNodeFocusBehavior = 0
	// SKNodeFocusBehaviorOccluding: Node is not focusable and prevents nodes that it visually obscures from becoming focusable.
	SKNodeFocusBehaviorOccluding SKNodeFocusBehavior = 1
)

func (e SKNodeFocusBehavior) String() string {
	switch e {
	case SKNodeFocusBehaviorFocusable:
		return "SKNodeFocusBehaviorFocusable"
	case SKNodeFocusBehaviorNone:
		return "SKNodeFocusBehaviorNone"
	case SKNodeFocusBehaviorOccluding:
		return "SKNodeFocusBehaviorOccluding"
	default:
		return fmt.Sprintf("SKNodeFocusBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKParticleRenderOrder
type SKParticleRenderOrder uint

const (
	// SKParticleRenderOrderDontCare: The particles can be rendered in any order.
	SKParticleRenderOrderDontCare SKParticleRenderOrder = 2
	// SKParticleRenderOrderOldestFirst: The particles are rendered from oldest to newest.
	SKParticleRenderOrderOldestFirst SKParticleRenderOrder = 1
	// SKParticleRenderOrderOldestLast: The particles are rendered from newest to oldest.
	SKParticleRenderOrderOldestLast SKParticleRenderOrder = 0
)

func (e SKParticleRenderOrder) String() string {
	switch e {
	case SKParticleRenderOrderDontCare:
		return "SKParticleRenderOrderDontCare"
	case SKParticleRenderOrderOldestFirst:
		return "SKParticleRenderOrderOldestFirst"
	case SKParticleRenderOrderOldestLast:
		return "SKParticleRenderOrderOldestLast"
	default:
		return fmt.Sprintf("SKParticleRenderOrder(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRepeatMode
type SKRepeatMode int

const (
	// SKRepeatModeClamp: When a sample is calculated, the time value is clamped to the range of time values found in the sequence.
	SKRepeatModeClamp SKRepeatMode = 1
	// SKRepeatModeLoop: When a sample is calculated, the sequence loops back to the beginning of the sequence.
	SKRepeatModeLoop SKRepeatMode = 2
)

func (e SKRepeatMode) String() string {
	switch e {
	case SKRepeatModeClamp:
		return "SKRepeatModeClamp"
	case SKRepeatModeLoop:
		return "SKRepeatModeLoop"
	default:
		return fmt.Sprintf("SKRepeatMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKSceneScaleMode
type SKSceneScaleMode int

const (
	// SKSceneScaleModeAspectFill: The scaling factor of each dimension is calculated and the larger of the two is chosen.
	SKSceneScaleModeAspectFill SKSceneScaleMode = 1
	// SKSceneScaleModeAspectFit: The scaling factor of each dimension is calculated and the smaller of the two is chosen.
	SKSceneScaleModeAspectFit SKSceneScaleMode = 2
	// SKSceneScaleModeFill: Each axis of the scene is scaled independently so that each axis in the scene exactly maps to the length of that axis in the view.
	SKSceneScaleModeFill SKSceneScaleMode = 0
	// SKSceneScaleModeResizeFill: The scene is not scaled to match the view.
	SKSceneScaleModeResizeFill SKSceneScaleMode = 3
)

func (e SKSceneScaleMode) String() string {
	switch e {
	case SKSceneScaleModeAspectFill:
		return "SKSceneScaleModeAspectFill"
	case SKSceneScaleModeAspectFit:
		return "SKSceneScaleModeAspectFit"
	case SKSceneScaleModeFill:
		return "SKSceneScaleModeFill"
	case SKSceneScaleModeResizeFill:
		return "SKSceneScaleModeResizeFill"
	default:
		return fmt.Sprintf("SKSceneScaleMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTextureFilteringMode
type SKTextureFilteringMode int

const (
	// SKTextureFilteringLinear: Each pixel is drawn by using a linear filter of multiple texels in the texture.
	SKTextureFilteringLinear SKTextureFilteringMode = 1
	// SKTextureFilteringNearest: Each pixel is drawn using the nearest point in the texture.
	SKTextureFilteringNearest SKTextureFilteringMode = 0
)

func (e SKTextureFilteringMode) String() string {
	switch e {
	case SKTextureFilteringLinear:
		return "SKTextureFilteringLinear"
	case SKTextureFilteringNearest:
		return "SKTextureFilteringNearest"
	default:
		return fmt.Sprintf("SKTextureFilteringMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTileAdjacencyMask
type SKTileAdjacencyMask uint

const (
	SKTileAdjacencyAll                 SKTileAdjacencyMask = 255
	SKTileAdjacencyDown                SKTileAdjacencyMask = 16
	SKTileAdjacencyDownEdge            SKTileAdjacencyMask = 199
	SKTileAdjacencyLeft                SKTileAdjacencyMask = 64
	SKTileAdjacencyLeftEdge            SKTileAdjacencyMask = 31
	SKTileAdjacencyLowerLeft           SKTileAdjacencyMask = 32
	SKTileAdjacencyLowerLeftCorner     SKTileAdjacencyMask = 253
	SKTileAdjacencyLowerLeftEdge       SKTileAdjacencyMask = 7
	SKTileAdjacencyLowerRight          SKTileAdjacencyMask = 8
	SKTileAdjacencyLowerRightCorner    SKTileAdjacencyMask = 127
	SKTileAdjacencyLowerRightEdge      SKTileAdjacencyMask = 193
	SKTileAdjacencyRight               SKTileAdjacencyMask = 4
	SKTileAdjacencyRightEdge           SKTileAdjacencyMask = 241
	SKTileAdjacencyUp                  SKTileAdjacencyMask = 1
	SKTileAdjacencyUpEdge              SKTileAdjacencyMask = 124
	SKTileAdjacencyUpperLeft           SKTileAdjacencyMask = 128
	SKTileAdjacencyUpperLeftCorner     SKTileAdjacencyMask = 247
	SKTileAdjacencyUpperLeftEdge       SKTileAdjacencyMask = 28
	SKTileAdjacencyUpperRight          SKTileAdjacencyMask = 2
	SKTileAdjacencyUpperRightCorner    SKTileAdjacencyMask = 223
	SKTileAdjacencyUpperRightEdge      SKTileAdjacencyMask = 112
	SKTileHexFlatAdjacencyAll          SKTileAdjacencyMask = 63
	SKTileHexFlatAdjacencyDown         SKTileAdjacencyMask = 8
	SKTileHexFlatAdjacencyLowerLeft    SKTileAdjacencyMask = 16
	SKTileHexFlatAdjacencyLowerRight   SKTileAdjacencyMask = 4
	SKTileHexFlatAdjacencyUp           SKTileAdjacencyMask = 1
	SKTileHexFlatAdjacencyUpperLeft    SKTileAdjacencyMask = 32
	SKTileHexFlatAdjacencyUpperRight   SKTileAdjacencyMask = 2
	SKTileHexPointyAdjacencyAdd        SKTileAdjacencyMask = 63
	SKTileHexPointyAdjacencyLeft       SKTileAdjacencyMask = 32
	SKTileHexPointyAdjacencyLowerLeft  SKTileAdjacencyMask = 16
	SKTileHexPointyAdjacencyLowerRight SKTileAdjacencyMask = 8
	SKTileHexPointyAdjacencyRight      SKTileAdjacencyMask = 4
	SKTileHexPointyAdjacencyUpperLeft  SKTileAdjacencyMask = 1
	SKTileHexPointyAdjacencyUpperRight SKTileAdjacencyMask = 2
)

func (e SKTileAdjacencyMask) String() string {
	switch e {
	case SKTileAdjacencyAll:
		return "SKTileAdjacencyAll"
	case SKTileAdjacencyDown:
		return "SKTileAdjacencyDown"
	case SKTileAdjacencyDownEdge:
		return "SKTileAdjacencyDownEdge"
	case SKTileAdjacencyLeft:
		return "SKTileAdjacencyLeft"
	case SKTileAdjacencyLeftEdge:
		return "SKTileAdjacencyLeftEdge"
	case SKTileAdjacencyLowerLeft:
		return "SKTileAdjacencyLowerLeft"
	case SKTileAdjacencyLowerLeftCorner:
		return "SKTileAdjacencyLowerLeftCorner"
	case SKTileAdjacencyLowerLeftEdge:
		return "SKTileAdjacencyLowerLeftEdge"
	case SKTileAdjacencyLowerRight:
		return "SKTileAdjacencyLowerRight"
	case SKTileAdjacencyLowerRightCorner:
		return "SKTileAdjacencyLowerRightCorner"
	case SKTileAdjacencyLowerRightEdge:
		return "SKTileAdjacencyLowerRightEdge"
	case SKTileAdjacencyRight:
		return "SKTileAdjacencyRight"
	case SKTileAdjacencyRightEdge:
		return "SKTileAdjacencyRightEdge"
	case SKTileAdjacencyUp:
		return "SKTileAdjacencyUp"
	case SKTileAdjacencyUpEdge:
		return "SKTileAdjacencyUpEdge"
	case SKTileAdjacencyUpperLeft:
		return "SKTileAdjacencyUpperLeft"
	case SKTileAdjacencyUpperLeftCorner:
		return "SKTileAdjacencyUpperLeftCorner"
	case SKTileAdjacencyUpperLeftEdge:
		return "SKTileAdjacencyUpperLeftEdge"
	case SKTileAdjacencyUpperRight:
		return "SKTileAdjacencyUpperRight"
	case SKTileAdjacencyUpperRightCorner:
		return "SKTileAdjacencyUpperRightCorner"
	case SKTileAdjacencyUpperRightEdge:
		return "SKTileAdjacencyUpperRightEdge"
	case SKTileHexFlatAdjacencyAll:
		return "SKTileHexFlatAdjacencyAll"
	default:
		return fmt.Sprintf("SKTileAdjacencyMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTileDefinitionRotation
type SKTileDefinitionRotation uint

const (
	SKTileDefinitionRotation0   SKTileDefinitionRotation = 0
	SKTileDefinitionRotation180 SKTileDefinitionRotation = 2
	SKTileDefinitionRotation270 SKTileDefinitionRotation = 3
	SKTileDefinitionRotation90  SKTileDefinitionRotation = 1
)

func (e SKTileDefinitionRotation) String() string {
	switch e {
	case SKTileDefinitionRotation0:
		return "SKTileDefinitionRotation0"
	case SKTileDefinitionRotation180:
		return "SKTileDefinitionRotation180"
	case SKTileDefinitionRotation270:
		return "SKTileDefinitionRotation270"
	case SKTileDefinitionRotation90:
		return "SKTileDefinitionRotation90"
	default:
		return fmt.Sprintf("SKTileDefinitionRotation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTileSetType
type SKTileSetType uint

const (
	SKTileSetTypeGrid            SKTileSetType = 0
	SKTileSetTypeHexagonalFlat   SKTileSetType = 2
	SKTileSetTypeHexagonalPointy SKTileSetType = 3
	SKTileSetTypeIsometric       SKTileSetType = 1
)

func (e SKTileSetType) String() string {
	switch e {
	case SKTileSetTypeGrid:
		return "SKTileSetTypeGrid"
	case SKTileSetTypeHexagonalFlat:
		return "SKTileSetTypeHexagonalFlat"
	case SKTileSetTypeHexagonalPointy:
		return "SKTileSetTypeHexagonalPointy"
	case SKTileSetTypeIsometric:
		return "SKTileSetTypeIsometric"
	default:
		return fmt.Sprintf("SKTileSetType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKTransitionDirection
type SKTransitionDirection int

const (
	// SKTransitionDirectionDown: The transition goes down.
	SKTransitionDirectionDown SKTransitionDirection = 1
	// SKTransitionDirectionLeft: The transition goes left.
	SKTransitionDirectionLeft SKTransitionDirection = 3
	// SKTransitionDirectionRight: The transition goes right.
	SKTransitionDirectionRight SKTransitionDirection = 2
	// SKTransitionDirectionUp: The transition goes up.
	SKTransitionDirectionUp SKTransitionDirection = 0
)

func (e SKTransitionDirection) String() string {
	switch e {
	case SKTransitionDirectionDown:
		return "SKTransitionDirectionDown"
	case SKTransitionDirectionLeft:
		return "SKTransitionDirectionLeft"
	case SKTransitionDirectionRight:
		return "SKTransitionDirectionRight"
	case SKTransitionDirectionUp:
		return "SKTransitionDirectionUp"
	default:
		return fmt.Sprintf("SKTransitionDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/SpriteKit/SKUniformType
type SKUniformType int

const (
	// SKUniformTypeFloat: Indicates that the uniform variable holds a 32-bit floating-point value.
	SKUniformTypeFloat SKUniformType = 1
	// SKUniformTypeFloatMatrix2: Indicates that the uniform variable holds a  matrix of four 32-bit floating-point values.
	SKUniformTypeFloatMatrix2 SKUniformType = 5
	// SKUniformTypeFloatMatrix3: Indicates that the uniform variable holds a  matrix of four 32-bit floating-point values.
	SKUniformTypeFloatMatrix3 SKUniformType = 6
	// SKUniformTypeFloatMatrix4: Indicates that the uniform variable holds a  matrix of four 32-bit floating-point values.
	SKUniformTypeFloatMatrix4 SKUniformType = 7
	// SKUniformTypeFloatVector2: Indicates that the uniform variable holds a vector of two 32-bit floating-point values.
	SKUniformTypeFloatVector2 SKUniformType = 2
	// SKUniformTypeFloatVector3: Indicates that the uniform variable holds a vector of three 32-bit floating-point values.
	SKUniformTypeFloatVector3 SKUniformType = 3
	// SKUniformTypeFloatVector4: Indicates that the uniform variable holds a vector of four 32-bit floating-point values.
	SKUniformTypeFloatVector4 SKUniformType = 4
	// SKUniformTypeNone: Indicates that the uniform variable does not currently hold any data.
	SKUniformTypeNone SKUniformType = 0
	// SKUniformTypeTexture: Indicates that the uniform variable holds a reference to a SpriteKit texture.
	SKUniformTypeTexture SKUniformType = 8
)

func (e SKUniformType) String() string {
	switch e {
	case SKUniformTypeFloat:
		return "SKUniformTypeFloat"
	case SKUniformTypeFloatMatrix2:
		return "SKUniformTypeFloatMatrix2"
	case SKUniformTypeFloatMatrix3:
		return "SKUniformTypeFloatMatrix3"
	case SKUniformTypeFloatMatrix4:
		return "SKUniformTypeFloatMatrix4"
	case SKUniformTypeFloatVector2:
		return "SKUniformTypeFloatVector2"
	case SKUniformTypeFloatVector3:
		return "SKUniformTypeFloatVector3"
	case SKUniformTypeFloatVector4:
		return "SKUniformTypeFloatVector4"
	case SKUniformTypeNone:
		return "SKUniformTypeNone"
	case SKUniformTypeTexture:
		return "SKUniformTypeTexture"
	default:
		return fmt.Sprintf("SKUniformType(%d)", e)
	}
}
