# MCPluginMaker
A Minecraft plugin maker using Maven, Go, and Java.
-----------------------------------------------------------
## Installation
* Download maven from https://maven.apache.org/download.cgi under the "Files" header
* Place extracted folder anywhere but keep note of it's location
* (Windows)
* In 'Run' Type 'Edit the system envirnment variables' and click on 'Environment Variables..' button near the bottom
* Double-Click on 'Path' in the 'User variables for {Name}' and Click on 'New' on the left near the top
* In the text area type '{Path to apache-maven folder}\bin' (I.E. 'C:\apache-maven-3.6.3\bin')
* Click 'OK', 'OK', 'OK'
* Download the provided .Exe from the above list and place in a folder on its own (It will create folders and a settings.json in that folder). The .exe will soon be hosted on a website with better instructions along with pictures (In development)
* Start using!
## Building Instructions
To build a project, simply hit the 'Build' button in the main toolbar (It's the save icon for right now) after you've made sure the project you are building has been selected from the list on the left. This will place a {Name}-1.0.jar file in your projects\{Name of project}\target folder. Just drag-and-drop that into your plugins folder. Plugin has been built and installed!
## Current Features
* Auto build
* Create/Load projects
* Create/Load commands
## Commands
Commands aren't fully there. You are able to build commands and 
the program will stop you from trying to create new commands with the same name/slashcommand in the GUI.
### Features to Come
1. Basic player commands (I.E /heal, /feed, /kill type commands)
2. Basic target player commands (I.E /heal {player}, feed {player}, kill {player})
3. Add permissions (optional)
## Crafting
Crafting hasn't been added yet, but will be an easy add when I get around to it.
### Features to Come
1. Crafting table crafting (Shaped and Unshaped)
2. Furnace crafting
3. Look into custom crafting with chests?
## Listeners
Listeners haven't been added as of yet. They will be added very shortly
### Features to Come
1. Block Listeners
2. Player Listeners
3. Entity Listeners
4. World Listeners
5. Damage Listeners
6. Cancellation for each listener
## Entities
Custom mobs will be makable using this GUI very easily. As they are very copy/pastable without too much work. 
Custom mob creation will start being worked on during the creation of listeners
### Features to Come
1. Custom names
2. Custom health, damage, speed
3. Custom drops
4. Ridable
5. Custom Goals/Targets
## Minigames
Minigames have always been something that I've loved to code. As such, I'm going to find a way to make it templatable and be makeable using this GUI.
### Features to Come
1. Arena Name
2. World creation type (Corners [Select corner 1, corner 2], New world [Would recommend using Paper MC])
3. Kits or not
4. Custom parent command (i.E /my_arena_cmd create, /my_arena_cmd join)
5. Sign support
6. Edit most attributes of a regular arena (max/min players, start timer, etc..)
