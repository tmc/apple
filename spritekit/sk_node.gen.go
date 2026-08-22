// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKNode] class.
var (
	_SKNodeClass     SKNodeClass
	_SKNodeClassOnce sync.Once
)

func getSKNodeClass() SKNodeClass {
	_SKNodeClassOnce.Do(func() {
		_SKNodeClass = SKNodeClass{class: objc.GetClass("SKNode")}
	})
	return _SKNodeClass
}

// GetSKNodeClass returns the class object for SKNode.
func GetSKNodeClass() SKNodeClass {
	return getSKNodeClass()
}

type SKNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKNodeClass) Alloc() SKNode {
	rv := objc.Send[SKNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The base class of all SpriteKit nodes.
//
// # Overview
//
// [SKNode] provides base properties for its subclasses and it can be used as
// a container or layout tool for other nodes. For example, you might add a
// collection of nodes as children to an [SKNode] that all move together
// within the scene; because nodes inherit the properties of their parent,
// changing the parent node’s [SKNode.Position] propagates the change to its
// children as well.
//
// [SKNode] does not draw any content itself. Its visual counterparts are
// listed in Nodes that Draw in [Nodes for Scene Building].
//
// # First Steps
//
//   - [SKNode.InitWithCoder]: Called when a node is initialized from an .sks file.
//
// # Positioning Content in a Scene
//
//   - [SKNode.Position]: The position of the node in its parent’s coordinate system.
//   - [SKNode.SetPosition]
//
// # Querying the Content Size
//
//   - [SKNode.Frame]: A rectangle in the parent’s coordinate system that contains the node’s content, ignoring the node’s children.
//   - [SKNode.CalculateAccumulatedFrame]: Returns a rectangle in the parent’s coordinate system that contains the position and size of itself and all child nodes.
//
// # Configuring Draw Order
//
//   - [SKNode.ZPosition]: The height of the node relative to its parent.
//   - [SKNode.SetZPosition]
//
// # Scaling and Rotating
//
//   - [SKNode.ZRotation]: The Euler rotation about the z axis (in radians).
//   - [SKNode.SetZRotation]
//   - [SKNode.SetScale]: Sets the [xScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/xScale>) and [yScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/yScale>) properties of the node.
//   - [SKNode.XScale]: A scaling factor that multiplies the width of a node and its children.
//   - [SKNode.SetXScale]
//   - [SKNode.YScale]: A scaling factor that multiplies the height of a node and its children.
//   - [SKNode.SetYScale]
//
// # Accessing Related Nodes
//
//   - [SKNode.Scene]: The scene node that contains this node.
//   - [SKNode.Parent]: The node’s parent node.
//   - [SKNode.Children]: The node’s children.
//
// # Modifying the Node Tree
//
//   - [SKNode.AddChild]: Adds a node to the end of the receiver’s list of child nodes.
//   - [SKNode.InsertChildAtIndex]: Inserts a node into a specific position in the receiver’s list of child nodes.
//   - [SKNode.IsEqualToNode]: Compares the parameter node to the receiving node.
//   - [SKNode.MoveToParent]: Moves the node to a new parent node in the scene.
//   - [SKNode.RemoveFromParent]: Removes the receiving node from its parent.
//   - [SKNode.RemoveAllChildren]: Removes all of the node’s children.
//   - [SKNode.RemoveChildrenInArray]: Removes a list of children from the receiving node.
//   - [SKNode.InParentHierarchy]: Returns a Boolean value that indicates whether the node is a descendant of the target node.
//
// # Accessing Nodes by Name
//
//   - [SKNode.Name]: The node’s assignable name.
//   - [SKNode.SetName]
//   - [SKNode.ChildNodeWithName]: Searches the children of the receiving node for a node with a specific name.
//   - [SKNode.EnumerateChildNodesWithNameUsingBlock]: Searches the children of the receiving node to perform processing for nodes that share a name.
//   - [SKNode.ObjectForKeyedSubscript]: Returns an array of nodes that match the name parameter.
//
// # Altering Node Visibility
//
//   - [SKNode.Alpha]: The transparency value applied to the node’s contents.
//   - [SKNode.SetAlpha]
//   - [SKNode.IsHidden]: A Boolean value that determines whether a node and its descendants are rendered.
//   - [SKNode.SetHidden]
//
// # Running Actions
//
//   - [SKNode.RunAction]: Adds an action to the list of actions executed by the node.
//   - [SKNode.RunActionCompletion]: Adds an action to the list of actions executed by the node and schedules the argument block to be run upon completion of the action.
//   - [SKNode.RunActionWithKey]: Adds an identifiable action to the list of actions executed by the node.
//   - [SKNode.Speed]: A speed modifier applied to all actions executed by a node and its descendants.
//   - [SKNode.SetSpeed]
//   - [SKNode.IsPaused]: A Boolean value that determines whether actions on the node and its descendants are processed.
//   - [SKNode.SetPaused]
//   - [SKNode.ActionForKey]: Returns an action associated with a specific key.
//   - [SKNode.HasActions]: Returns a Boolean value that indicates whether the node is executing actions.
//   - [SKNode.RemoveAllActions]: Ends and removes all actions from the node.
//   - [SKNode.RemoveActionForKey]: Removes an action associated with a specific key.
//
// # Adding Physics Behaviors
//
//   - [SKNode.PhysicsBody]: The physics body associated with the node.
//   - [SKNode.SetPhysicsBody]
//
// # Constraining Node Position or Rotation
//
//   - [SKNode.Constraints]: A list of constraints to apply to the node.
//   - [SKNode.SetConstraints]
//   - [SKNode.ReachConstraints]: The reach constraints to apply to the node when executing a reach action.
//   - [SKNode.SetReachConstraints]
//
// # Detecting Collisions Manually
//
//   - [SKNode.IntersectsNode]: Returns a Boolean value that indicates whether this node intersects the specified node.
//
// # Adding GameplayKit Behaviors
//
//   - [SKNode.Entity]: The GameplayKit entity this node represents.
//   - [SKNode.SetEntity]
//
// # Handling User Input
//
//   - [SKNode.IsUserInteractionEnabled]: A Boolean value that indicates whether the node receives touch events.
//   - [SKNode.SetUserInteractionEnabled]
//
// # Hit Testing
//
//   - [SKNode.ContainsPoint]: Returns a Boolean value that indicates whether a point lies inside the parent’s coordinate system.
//   - [SKNode.NodeAtPoint]: Returns the deepest visible descendant that intersects a point.
//   - [SKNode.NodesAtPoint]: Returns an array of all visible descendants that intersect a point.
//
// # Converting Between Coordinate Systems of Different Nodes
//
//   - [SKNode.ConvertPointFromNode]: Converts a point from the coordinate system of another node in the node tree to the coordinate system of this node.
//   - [SKNode.ConvertPointToNode]: Converts a point in this node’s coordinate system to the coordinate system of another node in the node tree.
//
// # Adding Custom Data Without Subclassing
//
//   - [SKNode.UserData]: A dictionary containing arbitrary data.
//   - [SKNode.SetUserData]
//
// # Providing Accessibility
//
//   - [SKNode.AccessibilityChildren]: An array of user interface elements that represent children of this element.
//   - [SKNode.SetAccessibilityChildren]
//   - [SKNode.AccessibilityFrame]: The size of this user interface element, in screen points.
//   - [SKNode.SetAccessibilityFrame]
//   - [SKNode.AccessibilityHelp]: The help description of this user interface element; for example, the text shown in a tooltip.
//   - [SKNode.SetAccessibilityHelp]
//   - [SKNode.AccessibilityLabel]: A short description of this user interface element.
//   - [SKNode.SetAccessibilityLabel]
//   - [SKNode.AccessibilityParent]: The user interface element that contains this element.
//   - [SKNode.SetAccessibilityParent]
//   - [SKNode.AccessibilityRole]: A string value describing the user interface element type; for example, a button.
//   - [SKNode.SetAccessibilityRole]
//   - [SKNode.AccessibilityRoleDescription]: A string value describing the user interface element name and type; for example, the Buy button.
//   - [SKNode.SetAccessibilityRoleDescription]
//   - [SKNode.AccessibilitySubrole]: A string that defines this user interface element’s subrole; for example, a full-screen button.
//   - [SKNode.SetAccessibilitySubrole]
//   - [SKNode.IsAccessibilityElement]: A toggle you implement to indicate to the system whether this user interface element should be exposed to the user.
//   - [SKNode.SetAccessibilityElement]
//   - [SKNode.IsAccessibilityEnabled]: A toggle you implement to indicate to the system whether this user interface element should respond to user input.
//   - [SKNode.SetAccessibilityEnabled]
//   - [SKNode.AccessibilityHitTest]: Returns the frontmost user interface element in the element hierarchy.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode
//
// [Nodes for Scene Building]: https://developer.apple.com/documentation/SpriteKit/nodes-for-scene-building
type SKNode struct {
	objectivec.Object
}

// SKNodeFromID constructs a [SKNode] from an objc.ID.
//
// The base class of all SpriteKit nodes.
func SKNodeFromID(id objc.ID) SKNode {
	return SKNode{objectivec.Object{ID: id}}
}

// NOTE: SKNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKNode] class.
//
// # First Steps
//
//   - [ISKNode.InitWithCoder]: Called when a node is initialized from an .sks file.
//
// # Positioning Content in a Scene
//
//   - [ISKNode.Position]: The position of the node in its parent’s coordinate system.
//   - [ISKNode.SetPosition]
//
// # Querying the Content Size
//
//   - [ISKNode.Frame]: A rectangle in the parent’s coordinate system that contains the node’s content, ignoring the node’s children.
//   - [ISKNode.CalculateAccumulatedFrame]: Returns a rectangle in the parent’s coordinate system that contains the position and size of itself and all child nodes.
//
// # Configuring Draw Order
//
//   - [ISKNode.ZPosition]: The height of the node relative to its parent.
//   - [ISKNode.SetZPosition]
//
// # Scaling and Rotating
//
//   - [ISKNode.ZRotation]: The Euler rotation about the z axis (in radians).
//   - [ISKNode.SetZRotation]
//   - [ISKNode.SetScale]: Sets the [xScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/xScale>) and [yScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/yScale>) properties of the node.
//   - [ISKNode.XScale]: A scaling factor that multiplies the width of a node and its children.
//   - [ISKNode.SetXScale]
//   - [ISKNode.YScale]: A scaling factor that multiplies the height of a node and its children.
//   - [ISKNode.SetYScale]
//
// # Accessing Related Nodes
//
//   - [ISKNode.Scene]: The scene node that contains this node.
//   - [ISKNode.Parent]: The node’s parent node.
//   - [ISKNode.Children]: The node’s children.
//
// # Modifying the Node Tree
//
//   - [ISKNode.AddChild]: Adds a node to the end of the receiver’s list of child nodes.
//   - [ISKNode.InsertChildAtIndex]: Inserts a node into a specific position in the receiver’s list of child nodes.
//   - [ISKNode.IsEqualToNode]: Compares the parameter node to the receiving node.
//   - [ISKNode.MoveToParent]: Moves the node to a new parent node in the scene.
//   - [ISKNode.RemoveFromParent]: Removes the receiving node from its parent.
//   - [ISKNode.RemoveAllChildren]: Removes all of the node’s children.
//   - [ISKNode.RemoveChildrenInArray]: Removes a list of children from the receiving node.
//   - [ISKNode.InParentHierarchy]: Returns a Boolean value that indicates whether the node is a descendant of the target node.
//
// # Accessing Nodes by Name
//
//   - [ISKNode.Name]: The node’s assignable name.
//   - [ISKNode.SetName]
//   - [ISKNode.ChildNodeWithName]: Searches the children of the receiving node for a node with a specific name.
//   - [ISKNode.EnumerateChildNodesWithNameUsingBlock]: Searches the children of the receiving node to perform processing for nodes that share a name.
//   - [ISKNode.ObjectForKeyedSubscript]: Returns an array of nodes that match the name parameter.
//
// # Altering Node Visibility
//
//   - [ISKNode.Alpha]: The transparency value applied to the node’s contents.
//   - [ISKNode.SetAlpha]
//   - [ISKNode.IsHidden]: A Boolean value that determines whether a node and its descendants are rendered.
//   - [ISKNode.SetHidden]
//
// # Running Actions
//
//   - [ISKNode.RunAction]: Adds an action to the list of actions executed by the node.
//   - [ISKNode.RunActionCompletion]: Adds an action to the list of actions executed by the node and schedules the argument block to be run upon completion of the action.
//   - [ISKNode.RunActionWithKey]: Adds an identifiable action to the list of actions executed by the node.
//   - [ISKNode.Speed]: A speed modifier applied to all actions executed by a node and its descendants.
//   - [ISKNode.SetSpeed]
//   - [ISKNode.IsPaused]: A Boolean value that determines whether actions on the node and its descendants are processed.
//   - [ISKNode.SetPaused]
//   - [ISKNode.ActionForKey]: Returns an action associated with a specific key.
//   - [ISKNode.HasActions]: Returns a Boolean value that indicates whether the node is executing actions.
//   - [ISKNode.RemoveAllActions]: Ends and removes all actions from the node.
//   - [ISKNode.RemoveActionForKey]: Removes an action associated with a specific key.
//
// # Adding Physics Behaviors
//
//   - [ISKNode.PhysicsBody]: The physics body associated with the node.
//   - [ISKNode.SetPhysicsBody]
//
// # Constraining Node Position or Rotation
//
//   - [ISKNode.Constraints]: A list of constraints to apply to the node.
//   - [ISKNode.SetConstraints]
//   - [ISKNode.ReachConstraints]: The reach constraints to apply to the node when executing a reach action.
//   - [ISKNode.SetReachConstraints]
//
// # Detecting Collisions Manually
//
//   - [ISKNode.IntersectsNode]: Returns a Boolean value that indicates whether this node intersects the specified node.
//
// # Adding GameplayKit Behaviors
//
//   - [ISKNode.Entity]: The GameplayKit entity this node represents.
//   - [ISKNode.SetEntity]
//
// # Handling User Input
//
//   - [ISKNode.IsUserInteractionEnabled]: A Boolean value that indicates whether the node receives touch events.
//   - [ISKNode.SetUserInteractionEnabled]
//
// # Hit Testing
//
//   - [ISKNode.ContainsPoint]: Returns a Boolean value that indicates whether a point lies inside the parent’s coordinate system.
//   - [ISKNode.NodeAtPoint]: Returns the deepest visible descendant that intersects a point.
//   - [ISKNode.NodesAtPoint]: Returns an array of all visible descendants that intersect a point.
//
// # Converting Between Coordinate Systems of Different Nodes
//
//   - [ISKNode.ConvertPointFromNode]: Converts a point from the coordinate system of another node in the node tree to the coordinate system of this node.
//   - [ISKNode.ConvertPointToNode]: Converts a point in this node’s coordinate system to the coordinate system of another node in the node tree.
//
// # Adding Custom Data Without Subclassing
//
//   - [ISKNode.UserData]: A dictionary containing arbitrary data.
//   - [ISKNode.SetUserData]
//
// # Providing Accessibility
//
//   - [ISKNode.AccessibilityChildren]: An array of user interface elements that represent children of this element.
//   - [ISKNode.SetAccessibilityChildren]
//   - [ISKNode.AccessibilityFrame]: The size of this user interface element, in screen points.
//   - [ISKNode.SetAccessibilityFrame]
//   - [ISKNode.AccessibilityHelp]: The help description of this user interface element; for example, the text shown in a tooltip.
//   - [ISKNode.SetAccessibilityHelp]
//   - [ISKNode.AccessibilityLabel]: A short description of this user interface element.
//   - [ISKNode.SetAccessibilityLabel]
//   - [ISKNode.AccessibilityParent]: The user interface element that contains this element.
//   - [ISKNode.SetAccessibilityParent]
//   - [ISKNode.AccessibilityRole]: A string value describing the user interface element type; for example, a button.
//   - [ISKNode.SetAccessibilityRole]
//   - [ISKNode.AccessibilityRoleDescription]: A string value describing the user interface element name and type; for example, the Buy button.
//   - [ISKNode.SetAccessibilityRoleDescription]
//   - [ISKNode.AccessibilitySubrole]: A string that defines this user interface element’s subrole; for example, a full-screen button.
//   - [ISKNode.SetAccessibilitySubrole]
//   - [ISKNode.IsAccessibilityElement]: A toggle you implement to indicate to the system whether this user interface element should be exposed to the user.
//   - [ISKNode.SetAccessibilityElement]
//   - [ISKNode.IsAccessibilityEnabled]: A toggle you implement to indicate to the system whether this user interface element should respond to user input.
//   - [ISKNode.SetAccessibilityEnabled]
//   - [ISKNode.AccessibilityHitTest]: Returns the frontmost user interface element in the element hierarchy.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode
type ISKNode interface {
	objectivec.IObject

	// Topic: First Steps

	// Called when a node is initialized from an .sks file.
	InitWithCoder(aDecoder foundation.INSCoder) SKNode

	// Topic: Positioning Content in a Scene

	// The position of the node in its parent’s coordinate system.
	Position() corefoundation.CGPoint
	SetPosition(value corefoundation.CGPoint)

	// Topic: Querying the Content Size

	// A rectangle in the parent’s coordinate system that contains the node’s content, ignoring the node’s children.
	Frame() corefoundation.CGRect
	// Returns a rectangle in the parent’s coordinate system that contains the position and size of itself and all child nodes.
	CalculateAccumulatedFrame() corefoundation.CGRect

	// Topic: Configuring Draw Order

	// The height of the node relative to its parent.
	ZPosition() float64
	SetZPosition(value float64)

	// Topic: Scaling and Rotating

	// The Euler rotation about the z axis (in radians).
	ZRotation() float64
	SetZRotation(value float64)
	// Sets the [xScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/xScale>) and [yScale](<https://developer.apple.com/documentation/SpriteKit/SKNode/yScale>) properties of the node.
	SetScale(scale float64)
	// A scaling factor that multiplies the width of a node and its children.
	XScale() float64
	SetXScale(value float64)
	// A scaling factor that multiplies the height of a node and its children.
	YScale() float64
	SetYScale(value float64)

	// Topic: Accessing Related Nodes

	// The scene node that contains this node.
	Scene() ISKScene
	// The node’s parent node.
	Parent() ISKNode
	// The node’s children.
	Children() []SKNode

	// Topic: Modifying the Node Tree

	// Adds a node to the end of the receiver’s list of child nodes.
	AddChild(node ISKNode)
	// Inserts a node into a specific position in the receiver’s list of child nodes.
	InsertChildAtIndex(node ISKNode, index int)
	// Compares the parameter node to the receiving node.
	IsEqualToNode(node ISKNode) bool
	// Moves the node to a new parent node in the scene.
	MoveToParent(parent ISKNode)
	// Removes the receiving node from its parent.
	RemoveFromParent()
	// Removes all of the node’s children.
	RemoveAllChildren()
	// Removes a list of children from the receiving node.
	RemoveChildrenInArray(nodes []SKNode)
	// Returns a Boolean value that indicates whether the node is a descendant of the target node.
	InParentHierarchy(parent ISKNode) bool

	// Topic: Accessing Nodes by Name

	// The node’s assignable name.
	Name() string
	SetName(value string)
	// Searches the children of the receiving node for a node with a specific name.
	ChildNodeWithName(name string) ISKNode
	// Searches the children of the receiving node to perform processing for nodes that share a name.
	EnumerateChildNodesWithNameUsingBlock(name string, block SKNodeBoolHandler)
	// Returns an array of nodes that match the name parameter.
	ObjectForKeyedSubscript(name string) []SKNode

	// Topic: Altering Node Visibility

	// The transparency value applied to the node’s contents.
	Alpha() float64
	SetAlpha(value float64)
	// A Boolean value that determines whether a node and its descendants are rendered.
	IsHidden() bool
	SetHidden(value bool)

	// Topic: Running Actions

	// Adds an action to the list of actions executed by the node.
	RunAction(action ISKAction)
	// Adds an action to the list of actions executed by the node and schedules the argument block to be run upon completion of the action.
	RunActionCompletion(action ISKAction, block VoidHandler)
	// Adds an identifiable action to the list of actions executed by the node.
	RunActionWithKey(action ISKAction, key string)
	// A speed modifier applied to all actions executed by a node and its descendants.
	Speed() float64
	SetSpeed(value float64)
	// A Boolean value that determines whether actions on the node and its descendants are processed.
	IsPaused() bool
	SetPaused(value bool)
	// Returns an action associated with a specific key.
	ActionForKey(key string) ISKAction
	// Returns a Boolean value that indicates whether the node is executing actions.
	HasActions() bool
	// Ends and removes all actions from the node.
	RemoveAllActions()
	// Removes an action associated with a specific key.
	RemoveActionForKey(key string)

	// Topic: Adding Physics Behaviors

	// The physics body associated with the node.
	PhysicsBody() ISKPhysicsBody
	SetPhysicsBody(value ISKPhysicsBody)

	// Topic: Constraining Node Position or Rotation

	// A list of constraints to apply to the node.
	Constraints() []SKConstraint
	SetConstraints(value []SKConstraint)
	// The reach constraints to apply to the node when executing a reach action.
	ReachConstraints() ISKReachConstraints
	SetReachConstraints(value ISKReachConstraints)

	// Topic: Detecting Collisions Manually

	// Returns a Boolean value that indicates whether this node intersects the specified node.
	IntersectsNode(node ISKNode) bool

	// Topic: Adding GameplayKit Behaviors

	// The GameplayKit entity this node represents.
	Entity() objectivec.IObject
	SetEntity(value objectivec.IObject)

	// Topic: Handling User Input

	// A Boolean value that indicates whether the node receives touch events.
	IsUserInteractionEnabled() bool
	SetUserInteractionEnabled(value bool)

	// Topic: Hit Testing

	// Returns a Boolean value that indicates whether a point lies inside the parent’s coordinate system.
	ContainsPoint(p corefoundation.CGPoint) bool
	// Returns the deepest visible descendant that intersects a point.
	NodeAtPoint(p corefoundation.CGPoint) ISKNode
	// Returns an array of all visible descendants that intersect a point.
	NodesAtPoint(p corefoundation.CGPoint) []SKNode

	// Topic: Converting Between Coordinate Systems of Different Nodes

	// Converts a point from the coordinate system of another node in the node tree to the coordinate system of this node.
	ConvertPointFromNode(point corefoundation.CGPoint, node ISKNode) corefoundation.CGPoint
	// Converts a point in this node’s coordinate system to the coordinate system of another node in the node tree.
	ConvertPointToNode(point corefoundation.CGPoint, node ISKNode) corefoundation.CGPoint

	// Topic: Adding Custom Data Without Subclassing

	// A dictionary containing arbitrary data.
	UserData() foundation.INSDictionary
	SetUserData(value foundation.INSDictionary)

	// Topic: Providing Accessibility

	// An array of user interface elements that represent children of this element.
	AccessibilityChildren() foundation.INSArray
	SetAccessibilityChildren(value foundation.INSArray)
	// The size of this user interface element, in screen points.
	AccessibilityFrame() corefoundation.CGRect
	SetAccessibilityFrame(value corefoundation.CGRect)
	// The help description of this user interface element; for example, the text shown in a tooltip.
	AccessibilityHelp() string
	SetAccessibilityHelp(value string)
	// A short description of this user interface element.
	AccessibilityLabel() string
	SetAccessibilityLabel(value string)
	// The user interface element that contains this element.
	AccessibilityParent() objectivec.IObject
	SetAccessibilityParent(value objectivec.IObject)
	// A string value describing the user interface element type; for example, a button.
	AccessibilityRole() string
	SetAccessibilityRole(value string)
	// A string value describing the user interface element name and type; for example, the Buy button.
	AccessibilityRoleDescription() string
	SetAccessibilityRoleDescription(value string)
	// A string that defines this user interface element’s subrole; for example, a full-screen button.
	AccessibilitySubrole() string
	SetAccessibilitySubrole(value string)
	// A toggle you implement to indicate to the system whether this user interface element should be exposed to the user.
	IsAccessibilityElement() bool
	SetAccessibilityElement(value bool)
	// A toggle you implement to indicate to the system whether this user interface element should respond to user input.
	IsAccessibilityEnabled() bool
	SetAccessibilityEnabled(value bool)
	// Returns the frontmost user interface element in the element hierarchy.
	AccessibilityHitTest(point corefoundation.CGPoint) objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (n SKNode) Init() SKNode {
	rv := objc.Send[SKNode](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n SKNode) Autorelease() SKNode {
	rv := objc.Send[SKNode](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKNode creates a new SKNode instance.
func NewSKNode() SKNode {
	class := getSKNodeClass()
	rv := objc.Send[SKNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewNodeWithCoder(aDecoder foundation.INSCoder) SKNode {
	instance := getSKNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKNodeFromID(rv)
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
func NewNodeWithFileNamed(filename string) SKNode {
	rv := objc.Send[objc.ID](objc.ID(getSKNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKNode{}, objc.ErrInitFailed
	}
	return SKNodeFromID(rv), nil
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func (n SKNode) InitWithCoder(aDecoder foundation.INSCoder) SKNode {
	rv := objc.Send[SKNode](n.ID, objc.Sel("initWithCoder:"), aDecoder)
	return rv
}

// Returns a rectangle in the parent’s coordinate system that contains the
// position and size of itself and all child nodes.
//
// # Discussion
//
// The frame takes into the account the cumulative effect of the
// [SKNode.XScale], [SKNode.YScale], and [SKNode.ZRotation] properties of each
// node in the subtree.
//
// Listing 1 shows, in Swift, how [SKNode.CalculateAccumulatedFrame] can be
// used display the bounding box of a shape node. The child node, although
// smaller than its parent, is rotated by 30° so that its bounds extend
// beyond its parent’s bounds. After `childNode` has been added to
// `parentNode`, a further shape node, `boundingBoxNode`, is created with its
// size based on the accumulated frame of parentNode.
//
// Listing 1. Displaying the accumulated frame of a shape node
//
// The figure below shows the result of Listing 1 with `parentNode` rendered
// in blue, `childNode` rendered in red and the `boundingBoxNode` rendered
// with a dashed line.
//
// [media-2793217]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/calculateAccumulatedFrame()
func (n SKNode) CalculateAccumulatedFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](n.ID, objc.Sel("calculateAccumulatedFrame"))
	return corefoundation.CGRect(rv)
}

// Sets the [SKNode.XScale] and [SKNode.YScale] properties of the node.
//
// scale: The new value to use for the node’s [SKNode.XScale] and [SKNode.YScale]
// properties.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/setScale(_:)
func (n SKNode) SetScale(scale float64) {
	objc.Send[objc.ID](n.ID, objc.Sel("setScale:"), scale)
}

// Adds a node to the end of the receiver’s list of child nodes.
//
// node: The node to add. The node must not already have a parent.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/addChild(_:)
func (n SKNode) AddChild(node ISKNode) {
	objc.Send[objc.ID](n.ID, objc.Sel("addChild:"), node)
}

// Inserts a node into a specific position in the receiver’s list of child
// nodes.
//
// node: The node to add. The node must not already have a parent.
//
// index: The position in the array to insert the node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/insertChild(_:at:)
func (n SKNode) InsertChildAtIndex(node ISKNode, index int) {
	objc.Send[objc.ID](n.ID, objc.Sel("insertChild:atIndex:"), node, index)
}

// Compares the parameter node to the receiving node.
//
// node: The node to compare to the receiving node.
//
// # Return Value
//
// true if the node is a descendant of the `parent` node; otherwise false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isEqual(to:)
func (n SKNode) IsEqualToNode(node ISKNode) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isEqualToNode:"), node)
	return rv
}

// Moves the node to a new parent node in the scene.
//
// parent: An [SKNode] object to move the receiver to. This node must be in the same
// scene as the node’s current parent.
//
// # Discussion
//
// The node maintains its current position in scene coordinates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/move(toParent:)
func (n SKNode) MoveToParent(parent ISKNode) {
	objc.Send[objc.ID](n.ID, objc.Sel("moveToParent:"), parent)
}

// Removes the receiving node from its parent.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/removeFromParent()
func (n SKNode) RemoveFromParent() {
	objc.Send[objc.ID](n.ID, objc.Sel("removeFromParent"))
}

// Removes all of the node’s children.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/removeAllChildren()
func (n SKNode) RemoveAllChildren() {
	objc.Send[objc.ID](n.ID, objc.Sel("removeAllChildren"))
}

// Removes a list of children from the receiving node.
//
// nodes: An array of [SKNode] objects that are all children of the receiving node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/removeChildren(in:)
func (n SKNode) RemoveChildrenInArray(nodes []SKNode) {
	objc.Send[objc.ID](n.ID, objc.Sel("removeChildrenInArray:"), objectivec.IObjectSliceToNSArray(nodes))
}

// Returns a Boolean value that indicates whether the node is a descendant of
// the target node.
//
// parent: An [SKNode] object to test against.
//
// # Return Value
//
// true if the node is a descendant of the `parent` node; otherwise false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/inParentHierarchy(_:)
func (n SKNode) InParentHierarchy(parent ISKNode) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("inParentHierarchy:"), parent)
	return rv
}

// Searches the children of the receiving node for a node with a specific
// name.
//
// name: The name to search for. This may be either the literal name of the node or
// a customized search string. See `Searching the Node Tree`.
//
// # Return Value
//
// If a node object with that name is found, the method returns the node
// object. Otherwise, it returns `nil`.
//
// # Discussion
//
// If more than one child share the same name, the first node discovered is
// returned.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/childNode(withName:)
func (n SKNode) ChildNodeWithName(name string) ISKNode {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("childNodeWithName:"), objc.String(name))
	return SKNodeFromID(rv)
}

// Searches the children of the receiving node to perform processing for nodes
// that share a name.
//
// name: The name to search for. This may be either the literal name of the node or
// a customized search string. See `Searching the Node Tree`.
//
// block: A block to execute on nodes that match the `name` parameter. The block has
// the signature `(` [SKNode] `, ` [UnsafeMutablePointer] `)`.
//
// # Discussion
//
// This method enumerates the child array in order, searching for nodes whose
// names match the search parameter. The block is called once for each node
// that matches the name parameter.
//
// The following Swift code shows how you could enumerate through the child
// nodes of a scene with a name containing the string `yellow`. Each matching
// node is hidden until the enumeration finds a node that also contains the
// string `triangle`. When this node is reached, `stop` is set to true and the
// processing stops.
//
// Listing 1. Enumerating child nodes
//
// You can also search by class name using
// [SKNode.EnumerateChildNodesWithNameUsingBlock]. However, for custom
// classes, you need to specify the fully annotated class name (i.e. the
// project name followed by the class name). The following Swift code shows a
// custom class, [SpaceshipNode], based on [SKSpriteNode], and created in a
// project named [SpaceGame]. The first search fails to return an instance of
// [SpaceshipNode] added as a child of `parentNode`:
//
// Listing 2. Enumerating child nodes
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/enumerateChildNodes(withName:using:)
//
// [UnsafeMutablePointer]: https://developer.apple.com/documentation/Swift/UnsafeMutablePointer
//
// [ObjCBool]: https://developer.apple.com/documentation/ObjectiveC/ObjCBool
func (n SKNode) EnumerateChildNodesWithNameUsingBlock(name string, block SKNodeBoolHandler) {
	_block1, _ := NewSKNodeBoolBlock(block)
	objc.Send[objc.ID](n.ID, objc.Sel("enumerateChildNodesWithName:usingBlock:"), objc.String(name), _block1)
}

// Returns an array of nodes that match the name parameter.
//
// name: The name to search for. This may be either the literal name of the node or
// a customized search string. See `Searching the Node Tree`.
//
// # Return Value
//
// An array of [SKNode] objects that match the name. If no matching nodes are
// found, an empty array is returned.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/subscript(_:)
func (n SKNode) ObjectForKeyedSubscript(name string) []SKNode {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(name))
	return objc.ConvertSlice(rv, func(id objc.ID) SKNode {
		return SKNodeFromID(id)
	})
}

// Adds an action to the list of actions executed by the node.
//
// action: The action to perform.
//
// # Discussion
//
// The new action is processed the next time the scene’s animation loop is
// processed.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/run(_:)
func (n SKNode) RunAction(action ISKAction) {
	objc.Send[objc.ID](n.ID, objc.Sel("runAction:"), action)
}

// Adds an action to the list of actions executed by the node and schedules
// the argument block to be run upon completion of the action.
//
// action: The action to perform.
//
// block: A completion block called when the action completes.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/run(_:completion:)
func (n SKNode) RunActionCompletion(action ISKAction, block VoidHandler) {
	_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](n.ID, objc.Sel("runAction:completion:"), action, _block1)
}

// Adds an identifiable action to the list of actions executed by the node.
//
// action: The action to perform.
//
// key: A unique key used to identify the action.
//
// # Discussion
//
// This method is identical to [SKNode.RunAction], but the action is stored so
// that it can be retrieved later. If an action using the same key is already
// running, it is removed before the new action is added.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/run(_:withKey:)
func (n SKNode) RunActionWithKey(action ISKAction, key string) {
	objc.Send[objc.ID](n.ID, objc.Sel("runAction:withKey:"), action, objc.String(key))
}

// Returns an action associated with a specific key.
//
// key: A string that uniquely identifies an action.
//
// # Return Value
//
// If an action exists that matches the key, the action object is returned.
// Otherwise, `nil` is returned.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/action(forKey:)
func (n SKNode) ActionForKey(key string) ISKAction {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("actionForKey:"), objc.String(key))
	return SKActionFromID(rv)
}

// Returns a Boolean value that indicates whether the node is executing
// actions.
//
// # Return Value
//
// true if the node has any executing actions; otherwise false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/hasActions()
func (n SKNode) HasActions() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasActions"))
	return rv
}

// Ends and removes all actions from the node.
//
// # Discussion
//
// When an action is removed from the node, any remaining animation the action
// would perform is skipped; however, previous changes are not reverted. It is
// possible that an action may make a final change to the scene when removed;
// if so, it is documented for the specific action in [SKAction].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/removeAllActions()
func (n SKNode) RemoveAllActions() {
	objc.Send[objc.ID](n.ID, objc.Sel("removeAllActions"))
}

// Removes an action associated with a specific key.
//
// key: A string that uniquely identifies an action.
//
// # Discussion
//
// If an action is found that matches the key, it is removed from the node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/removeAction(forKey:)
func (n SKNode) RemoveActionForKey(key string) {
	objc.Send[objc.ID](n.ID, objc.Sel("removeActionForKey:"), objc.String(key))
}

// Returns a Boolean value that indicates whether this node intersects the
// specified node.
//
// node: Another node in the same node tree.
//
// # Return Value
//
// true if the two nodes intersect; otherwise false.
//
// # Discussion
//
// The two nodes are considered to intersect if their frames intersect. The
// children of both nodes are ignored in this test.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/intersects(_:)
func (n SKNode) IntersectsNode(node ISKNode) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("intersectsNode:"), node)
	return rv
}

// Returns a Boolean value that indicates whether a point lies inside the
// parent’s coordinate system.
//
// p: A [CGPoint] to test against.
//
// # Return Value
//
// true if the point lies inside the parent’s coordinate system; otherwise
// false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/contains(_:)
func (n SKNode) ContainsPoint(p corefoundation.CGPoint) bool {
	rv := objc.Send[bool](n.ID, objc.Sel("containsPoint:"), p)
	return rv
}

// Returns the deepest visible descendant that intersects a point.
//
// p: A point in the node’s coordinate system.
//
// # Return Value
//
// A descendant in the subtree that intersects the point, or the receiver if
// no nodes intersect the point. Only nodes that have an [SKNode.Hidden] of
// `false` and an [SKNode.Alpha] greater that zero are returned. If multiple
// descendants intersect the point, the deepest node in the tree is returned.
// If multiple nodes are at the same level, the intersecting node with the
// largest z position is returned.
//
// # Discussion
//
// A point is considered to be in a node if it lies inside the rectangle
// returned by the [SKNode.CalculateAccumulatedFrame] method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/atPoint(_:)
func (n SKNode) NodeAtPoint(p corefoundation.CGPoint) ISKNode {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("nodeAtPoint:"), p)
	return SKNodeFromID(rv)
}

// Returns an array of all visible descendants that intersect a point.
//
// p: A point in the node’s coordinate system.
//
// # Return Value
//
// An array of all [SKNode] objects in the subtree that intersect the point.
// Only nodes that have an [SKNode.Hidden] of `false` and an [SKNode.Alpha]
// greater that zero are included in the returned array. If no nodes intersect
// the point, an empty array is returned.
//
// # Discussion
//
// A point is considered to be in a node if it lies inside the rectangle
// returned by the [SKNode.CalculateAccumulatedFrame] method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/nodes(at:)
func (n SKNode) NodesAtPoint(p corefoundation.CGPoint) []SKNode {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("nodesAtPoint:"), p)
	return objc.ConvertSlice(rv, func(id objc.ID) SKNode {
		return SKNodeFromID(id)
	})
}

// Converts a point from the coordinate system of another node in the node
// tree to the coordinate system of this node.
//
// point: A point in the other node’s coordinate system.
//
// node: Another node in the same node tree as this node.
//
// # Return Value
//
// The same point converted to this node’s coordinate system.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/convert(_:from:)
func (n SKNode) ConvertPointFromNode(point corefoundation.CGPoint, node ISKNode) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](n.ID, objc.Sel("convertPoint:fromNode:"), point, node)
	return corefoundation.CGPoint(rv)
}

// Converts a point in this node’s coordinate system to the coordinate
// system of another node in the node tree.
//
// point: A point in this node’s coordinate system.
//
// node: Another node in the same node tree as this node.
//
// # Return Value
//
// The same point converted to the other node’s coordinate system.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/convert(_:to:)
func (n SKNode) ConvertPointToNode(point corefoundation.CGPoint, node ISKNode) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](n.ID, objc.Sel("convertPoint:toNode:"), point, node)
	return corefoundation.CGPoint(rv)
}

// Returns the frontmost user interface element in the element hierarchy.
//
// point: Relative to the bottom-left of the screen, in screen points, and guaranteed
// to lie within the receiver.
//
// # Discussion
//
// Override this method to implement your own, deeper hit testing within a
// user interface element.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityHitTest(_:)
func (n SKNode) AccessibilityHitTest(point corefoundation.CGPoint) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityHitTest:"), point)
	return objectivec.Object{ID: rv}
}
func (n SKNode) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](n.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Converts each node into an obstacle by transforming its bounds into the
// scene’s coordinate system.
//
// nodes: An array of [SKNode] objects.
//
// # Return Value
//
// An array of [GKPolygonObstacle] objects.
//
// # Discussion
//
// Use the array of obstacles to create an obstacle graph ([GKObstacleGraph])
// in GameplayKit. See [GameplayKit] and [GameplayKit Programming Guide].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromNodeBounds:)
//
// [GKPolygonObstacle]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle
// [GKObstacleGraph]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph
// [GameplayKit Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/GameplayKit_Guide/index.html#//apple_ref/doc/uid/TP40015172
// [GameplayKit]: https://developer.apple.com/documentation/GameplayKit
func (_SKNodeClass SKNodeClass) ObstaclesFromNodeBounds(nodes []SKNode) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](objc.ID(_SKNodeClass.class), objc.Sel("obstaclesFromNodeBounds:"), objectivec.IObjectSliceToNSArray(nodes))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Converts each node into an obstacle by transforming the node’s physics
// body shape into the scene’s coordinate system.
//
// nodes: An array of [SKNode] objects.
//
// # Return Value
//
// An array of [GKPolygonObstacle] objects.
//
// # Discussion
//
// Use the array of obstacles to create an obstacle graph ([GKObstacleGraph])
// in GameplayKit. See [GameplayKit] and [GameplayKit Programming Guide].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromNodePhysicsBodies:)
//
// [GKPolygonObstacle]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle
// [GKObstacleGraph]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph
// [GameplayKit Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/GameplayKit_Guide/index.html#//apple_ref/doc/uid/TP40015172
// [GameplayKit]: https://developer.apple.com/documentation/GameplayKit
func (_SKNodeClass SKNodeClass) ObstaclesFromNodePhysicsBodies(nodes []SKNode) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](objc.ID(_SKNodeClass.class), objc.Sel("obstaclesFromNodePhysicsBodies:"), objectivec.IObjectSliceToNSArray(nodes))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Turns each node into an obstacle by changing the node’s texture into a
// physics shape and converting it into the scene’s coordinate system.
//
// sprites: An array of [SKNode] objects.
//
// accuracy: A floating point value between `0.001` and `1.0`, inclusive. Higher values
// create a more precise (but more complex) representation of the obstacle.
//
// # Return Value
//
// An array of [GKPolygonObstacle] objects.
//
// # Discussion
//
// Use the array of obstacles to create an obstacle graph ([GKObstacleGraph])
// in GameplayKit. See [GameplayKit] and [GameplayKit Programming Guide].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/obstacles(fromSpriteTextures:accuracy:)
//
// [GKPolygonObstacle]: https://developer.apple.com/documentation/GameplayKit/GKPolygonObstacle
// [GKObstacleGraph]: https://developer.apple.com/documentation/GameplayKit/GKObstacleGraph
// [GameplayKit Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/GameplayKit_Guide/index.html#//apple_ref/doc/uid/TP40015172
// [GameplayKit]: https://developer.apple.com/documentation/GameplayKit
func (_SKNodeClass SKNodeClass) ObstaclesFromSpriteTexturesAccuracy(sprites []SKNode, accuracy float32) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](objc.ID(_SKNodeClass.class), objc.Sel("obstaclesFromSpriteTextures:accuracy:"), objectivec.IObjectSliceToNSArray(sprites), accuracy)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The position of the node in its parent’s coordinate system.
//
// # Discussion
//
// The default value is `(0.0,0.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/position
func (n SKNode) Position() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](n.ID, objc.Sel("position"))
	return corefoundation.CGPoint(rv)
}
func (n SKNode) SetPosition(value corefoundation.CGPoint) {
	objc.Send[struct{}](n.ID, objc.Sel("setPosition:"), value)
}

// A rectangle in the parent’s coordinate system that contains the node’s
// content, ignoring the node’s children.
//
// # Discussion
//
// The frame is the smallest rectangle that contains the node’s content,
// taking into account the node’s [SKNode.XScale], [SKNode.YScale], and
// [SKNode.ZRotation] properties.
//
// Since [SKNode] does not draw content of its own, its frame size is
// arbitrary; it’s the visual subclasses of [SKNode] that do draw that
// define the frame’s size with a meaningful value that encloses its visual
// content.
//
// To get a rect that encloses all the child nodes of an [SKNode] parent
// object, use [SKNode.CalculateAccumulatedFrame].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/frame
func (n SKNode) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](n.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}

// The height of the node relative to its parent.
//
// # Discussion
//
// The default value is `0.0`. The positive z axis is projected toward the
// viewer so that nodes with larger z-position values are closer to the
// viewer. When a node tree is rendered, the height of each node (in absolute
// coordinates) is calculated and then all nodes in the tree are rendered from
// smallest z-position value to largest z-position value. If multiple nodes
// share the same z-position, those nodes are sorted so that parent nodes are
// drawn before their children, and siblings are rendered in the order that
// they appear in their parent’s [SKNode.Children] array. Hit-testing is
// processed in the opposite order.
//
// The [SKView] class’s [SKView.IgnoresSiblingOrder] property controls
// whether node sorting is enabled for nodes at the same z-position.
//
// # Using a Node’s Depth to Add Effects
//
// SpriteKit uses the [SKNode.ZPosition] value only to determine the hit
// testing and drawing order. You can also the z position to implement your
// own game effects. For example, you might use the height of a node to
// determine how it is rendered or how it moves onscreen. In this way, you can
// simulate fog or parallax effects. SpriteKit does not create these effects
// for you. Usually, you implement them by processing the scene immediately
// before it is rendered.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/zPosition
func (n SKNode) ZPosition() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("zPosition"))
	return rv
}
func (n SKNode) SetZPosition(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setZPosition:"), value)
}

// The Euler rotation about the z axis (in radians).
//
// # Discussion
//
// The default value is `0.0`, which indicates no rotation. A positive value
// indicates a counterclockwise rotation. When the coordinate system is
// rotated, it affects the node and its descendants. The rotation affects the
// node’s [SKNode.Frame] property, hit testing, rendering, and other similar
// characteristics.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/zRotation
func (n SKNode) ZRotation() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("zRotation"))
	return rv
}
func (n SKNode) SetZRotation(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setZRotation:"), value)
}

// A scaling factor that multiplies the width of a node and its children.
//
// # Discussion
//
// The [SKNode.XScale] property scales the width of the node and all of its
// descendants. The scale value affects how a node’s frame is calculated,
// its hit test area, how it is drawn, and other similar characteristics. The
// default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/xScale
func (n SKNode) XScale() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("xScale"))
	return rv
}
func (n SKNode) SetXScale(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setXScale:"), value)
}

// A scaling factor that multiplies the height of a node and its children.
//
// # Discussion
//
// The [SKNode.YScale] property scales the height of the node and all of its
// descendants. The scale value affects how a node’s frame is calculated,
// its hit test area, how it is drawn, and other similar characteristics. The
// default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/yScale
func (n SKNode) YScale() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("yScale"))
	return rv
}
func (n SKNode) SetYScale(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setYScale:"), value)
}

// The scene node that contains this node.
//
// # Discussion
//
// If the node is not embedded in a scene, the value is `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/scene
func (n SKNode) Scene() ISKScene {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("scene"))
	return SKSceneFromID(objc.ID(rv))
}

// The node’s parent node.
//
// # Discussion
//
// If the node is not in a node tree, the value is `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/parent
func (n SKNode) Parent() ISKNode {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parent"))
	return SKNodeFromID(objc.ID(rv))
}

// The node’s children.
//
// # Discussion
//
// The objects in this array are all [SKNode] objects.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/children
func (n SKNode) Children() []SKNode {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("children"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKNode {
		return SKNodeFromID(id)
	})
}

// The node’s assignable name.
//
// # Discussion
//
// This property is used to identify a node in other parts of your game logic.
// For example, you might use this name as part of collision testing. You can
// also search for nodes in a tree by their name.
//
// When choosing a name for a node, decide whether each node gets a unique
// name or whether some nodes will share a common name. If you give the node a
// unique name, you can find the node later by calling the
// [SKNode.ChildNodeWithName] method. If a name is shared by multiple nodes,
// the name usually means that these are all a similar object type in your
// game. In this case, you can iterate over those objects by calling the
// [SKNode.EnumerateChildNodesWithNameUsingBlock] method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/name
func (n SKNode) Name() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setName:"), objc.String(value))
}

// The transparency value applied to the node’s contents.
//
// # Discussion
//
// The default value is `1.0`.
//
// The [SKNode] class does not perform drawing, but many of its subclasses do.
// When a node or any of its descendants are drawn, the alpha component of
// each pixel is multiplied by the node’s [SKNode.Alpha] property and then
// clamped to the range `0.0`-`1.0`. This modified alpha value is used to
// blend the pixel into the framebuffer. Subclasses that render content define
// properties that determine the blending operations used in conjunction with
// the alpha value to blend pixels into the parent’s framebuffer.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/alpha
func (n SKNode) Alpha() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("alpha"))
	return rv
}
func (n SKNode) SetAlpha(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setAlpha:"), value)
}

// A Boolean value that determines whether a node and its descendants are
// rendered.
//
// # Discussion
//
// When hidden, a node and its descendants are not rendered. However, they
// still exist in the scene and continue to interact in other ways. For
// example, the node’s actions still run and the node can still be
// intersected with other nodes. The default value is false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isHidden
func (n SKNode) IsHidden() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isHidden"))
	return rv
}
func (n SKNode) SetHidden(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHidden:"), value)
}

// A speed modifier applied to all actions executed by a node and its
// descendants.
//
// # Discussion
//
// The default value is `1.0`, which means that all actions run at their
// normal speed. If you set a different speed, time appears to run faster or
// slower for all actions executed on the node and its descendants. For
// example, if you set a speed value of `2.0`, actions run twice as fast.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/speed
func (n SKNode) Speed() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("speed"))
	return rv
}
func (n SKNode) SetSpeed(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setSpeed:"), value)
}

// A Boolean value that determines whether actions on the node and its
// descendants are processed.
//
// # Discussion
//
// If the value is true, the node (and all of its descendants) are skipped
// when a scene processes actions.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isPaused
func (n SKNode) IsPaused() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isPaused"))
	return rv
}
func (n SKNode) SetPaused(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setPaused:"), value)
}

// The physics body associated with the node.
//
// # Discussion
//
// The default value is `nil`, which indicates that the node does not
// participate in the physics simulation at all. If a physics body is
// provided, when the scene’s physics are simulated, the physics body
// updates the node’s position and rotates the node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/physicsBody
func (n SKNode) PhysicsBody() ISKPhysicsBody {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("physicsBody"))
	return SKPhysicsBodyFromID(objc.ID(rv))
}
func (n SKNode) SetPhysicsBody(value ISKPhysicsBody) {
	objc.Send[struct{}](n.ID, objc.Sel("setPhysicsBody:"), value)
}

// A list of constraints to apply to the node.
//
// # Discussion
//
// Assign an array of [SKConstraint] objects to the node. The scene processes
// these constraints before the scene is rendered. The constraints are
// processed in array order. If multiple nodes in the node tree have
// constraints, there is no guaranteed order that the nodes are processed in.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/constraints
func (n SKNode) Constraints() []SKConstraint {
	rv := objc.Send[[]objc.ID](n.ID, objc.Sel("constraints"))
	return objc.ConvertSlice(rv, func(id objc.ID) SKConstraint {
		return SKConstraintFromID(id)
	})
}
func (n SKNode) SetConstraints(value []SKConstraint) {
	objc.Send[struct{}](n.ID, objc.Sel("setConstraints:"), objectivec.IObjectSliceToNSArray(value))
}

// The reach constraints to apply to the node when executing a reach action.
//
// # Discussion
//
// To use inverse kinematics, create a new [SKReachConstraints] object and
// assign it to this property. When a reach action calculates the new
// positions of this node, the possible values for this node are restricted to
// the constraints defined by this object. For more information on the inverse
// kinematic actions, see [SKAction].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/reachConstraints
func (n SKNode) ReachConstraints() ISKReachConstraints {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("reachConstraints"))
	return SKReachConstraintsFromID(objc.ID(rv))
}
func (n SKNode) SetReachConstraints(value ISKReachConstraints) {
	objc.Send[struct{}](n.ID, objc.Sel("setReachConstraints:"), value)
}

// The GameplayKit entity this node represents.
//
// # Discussion
//
// The Entity-Component architecture in the GameplayKit framework is a way to
// more easily manage complex object graphs in your game. For more information
// on this architecture, read [Entities and Components] in [GameplayKit
// Programming Guide].
//
// When you add entities (and their components) to a scene in the Xcode
// SpriteKit Scene Editor, Xcode automatically archive those entities
// alongside the SpriteKit scene content. Use the [GKScene] class to load the
// SpriteKit scene with its associated GameplayKit objects. Each entity
// associated with a SpriteKit node has a [GKSKNodeComponent] object that
// manages the relationship between the node and the [GKEntity] object it
// represents.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/entity
//
// [Entities and Components]: https://developer.apple.com/library/archive/documentation/General/Conceptual/GameplayKit_Guide/EntityComponent.html#//apple_ref/doc/uid/TP40015172-CH6
// [GKEntity]: https://developer.apple.com/documentation/GameplayKit/GKEntity
// [GKSKNodeComponent]: https://developer.apple.com/documentation/GameplayKit/GKSKNodeComponent
// [GKScene]: https://developer.apple.com/documentation/GameplayKit/GKScene
// [GameplayKit Programming Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/GameplayKit_Guide/index.html#//apple_ref/doc/uid/TP40015172
func (n SKNode) Entity() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("entity"))
	return objectivec.Object{ID: rv}
}
func (n SKNode) SetEntity(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setEntity:"), value)
}

// A Boolean value that indicates whether the node receives touch events.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isUserInteractionEnabled
func (n SKNode) IsUserInteractionEnabled() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isUserInteractionEnabled"))
	return rv
}
func (n SKNode) SetUserInteractionEnabled(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setUserInteractionEnabled:"), value)
}

// A dictionary containing arbitrary data.
//
// # Discussion
//
// You use this property to store your own data in a node. For example, you
// might store game-specific data about each node to use inside your game
// logic. This can be a useful alternative to creating your own node
// subclasses to hold game data.
//
// SpriteKit does not do anything with the data stored in the node. However,
// the data is archived when the node is archived.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/userData
func (n SKNode) UserData() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("userData"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n SKNode) SetUserData(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setUserData:"), value)
}

// An array of user interface elements that represent children of this
// element.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityChildren
func (n SKNode) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n SKNode) SetAccessibilityChildren(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityChildren:"), value)
}

// The size of this user interface element, in screen points.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityFrame
func (n SKNode) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](n.ID, objc.Sel("accessibilityFrame"))
	return corefoundation.CGRect(rv)
}
func (n SKNode) SetAccessibilityFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityFrame:"), value)
}

// The help description of this user interface element; for example, the text
// shown in a tooltip.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityHelp
func (n SKNode) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetAccessibilityHelp(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityHelp:"), objc.String(value))
}

// A short description of this user interface element.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityLabel
func (n SKNode) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetAccessibilityLabel(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityLabel:"), objc.String(value))
}

// The user interface element that contains this element.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityParent
func (n SKNode) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}
func (n SKNode) SetAccessibilityParent(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityParent:"), value)
}

// A string value describing the user interface element type; for example, a
// button.
//
// # Discussion
//
// See [Roles] for more information.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityRole
//
// [Roles]: https://developer.apple.com/documentation/applicationservices/carbon_accessibility/roles
func (n SKNode) AccessibilityRole() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityRole"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetAccessibilityRole(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityRole:"), objc.String(value))
}

// A string value describing the user interface element name and type; for
// example, the Buy button.
//
// # Discussion
//
// See [Roles] for more information.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilityRoleDescription
//
// [Roles]: https://developer.apple.com/documentation/applicationservices/carbon_accessibility/roles
func (n SKNode) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetAccessibilityRoleDescription(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(value))
}

// A string that defines this user interface element’s subrole; for example,
// a full-screen button.
//
// # Discussion
//
// See [Subroles] for more information.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/accessibilitySubrole
func (n SKNode) AccessibilitySubrole() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("accessibilitySubrole"))
	return foundation.NSStringFromID(rv).String()
}
func (n SKNode) SetAccessibilitySubrole(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(value))
}

// A toggle you implement to indicate to the system whether this user
// interface element should be exposed to the user.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isAccessibilityElement
func (n SKNode) IsAccessibilityElement() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isAccessibilityElement"))
	return rv
}
func (n SKNode) SetAccessibilityElement(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityElement:"), value)
}

// A toggle you implement to indicate to the system whether this user
// interface element should respond to user input.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/isAccessibilityEnabled
func (n SKNode) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}
func (n SKNode) SetAccessibilityEnabled(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setAccessibilityEnabled:"), value)
}

// RunActionCompletionSync is a synchronous wrapper around [SKNode.RunActionCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (n SKNode) RunActionCompletionSync(ctx context.Context, action ISKAction) error {
	done := make(chan struct{}, 1)
	n.RunActionCompletion(action, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
