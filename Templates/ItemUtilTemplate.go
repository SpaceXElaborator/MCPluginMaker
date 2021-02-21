package PluginTemplates

func GetItemUtilClassTemplate() string {
	return `package com.{{.Author}}.net;

import java.util.List;

public class Item {
	
	private String Name;
	private String Material;
	private List<String> Lore;
	
	public Item(String name, String mat, List<String> lore) {
		Name = name;
		Material = mat;
		Lore = lore;
	}
	
	public String getName() {
		return Name;
	}
	
	public String getMaterial() {
		return Material;
	}
	
	public List<String> getLore() {
		return Lore;
	}
	
}`
}