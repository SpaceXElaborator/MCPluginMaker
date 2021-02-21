package PluginTemplates

func GetItemTemplate() string {
	return `package com.{{.Author}}.net;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.bukkit.Material;
import org.bukkit.inventory.ItemStack;
import org.bukkit.inventory.meta.ItemMeta;

public class {{.Name}}CustomItems {

	private static List<Item> items = new ArrayList<Item>();
	
	public static void items() {
		{{with .Items -}}{{range $val := .}}
		// <<ITEM:{{$val.ItemMaterial}}||{{$val.ItemName}}||{{with $val.ItemDescription -}}{{range $strings := .}}{{.}}\n{{end}}{{end}}>>
		RegisterItem("{{$val.ItemName}}", "{{$val.ItemMaterial}}", Arrays.asList({{with $val.ItemDescription -}}{{range $ind, $strings := .}}{{if $ind}},{{end}}"{{.}}"{{end}}{{end}}));
		{{end}}{{end}}
	}
	
	private static void RegisterItem(String itemName, String mat, List<String> lore) {
		items.add(new Item(itemName, mat, lore));
	}
	
	public static ItemStack build(String name, Integer amount) {
		Item itm = null;
		for(Item i : items) {
			if(i.getName().equalsIgnoreCase(name)) {
				itm = i;
			}
		}
		ItemStack item = new ItemStack(Material.valueOf(itm.getMaterial()), amount);
		ItemMeta meta = item.getItemMeta();
		meta.setDisplayName(itm.getName());
		meta.setLore(itm.getLore());
		item.setItemMeta(meta);
		return item;
	}
	
}`
}