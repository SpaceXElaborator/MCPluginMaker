package pluginitems

// CustomItem Holds the material, name, and list of strings to be applied to the item on creation
type CustomItem struct {
	ItemMaterial, ItemName string
	ItemDescription        []string
}
