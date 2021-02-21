package PluginTemplates

func GetItemTemplate() string {
	return `package com.{{.Author}}.net;

import java.util.ArrayList;
import java.util.List;

import org.bukkit.Material;
import org.bukkit.inventory.ItemStack;
import org.bukkit.inventory.meta.ItemMeta;

public class {{.Name}}CustomItems {
	public static List<ItemStack> items = new ArrayList<ItemStack>();
	
	public static void RegisterItems() {
		{{with .Items -}}{{range $val := .}}
		// <<ITEM:{{$val.ItemMaterial}}||{{$val.ItemName}}||{{with $val.ItemDescription -}}{{range $strings := .}}{{.}}\n{{end}}{{end}}>>
		ItemStack {{$val.ItemName}}Item = new ItemStack(Material.valueOf("{{$val.ItemMaterial}}"));
		ItemMeta {{$val.ItemName}}Meta = {{$val.ItemName}}Item.getItemMeta();
		{{$val.ItemName}}Meta.setDisplayName("{{$val.ItemName}}");
		List<String> {{$val.ItemName}}Lore = new ArrayList<String>();
		{{with $val.ItemDescription -}}{{range $strings := .}}{{$val.ItemName}}Lore.add("{{.}}");
		{{end}}{{end}}{{$val.ItemName}}Meta.setLore({{$val.ItemName}}Lore);
		{{$val.ItemName}}Item.setItemMeta({{$val.ItemName}}Meta);
		items.add({{$val.ItemName}}Item);
		{{end}}{{end}}
	}
	
	public static ItemStack getItem(String name) {
		for(ItemStack item : items) {
			if(item.getItemMeta().getDisplayName().equalsIgnoreCase(name)) {
				return item;
			}
		}
		return null;
	}
	
}
`
}