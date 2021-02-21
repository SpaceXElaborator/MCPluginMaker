package PluginTemplates

func GetPomTemplate() string {
	return `<project xmlns="http://maven.apache.org/POM/4.0.0"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 https://maven.apache.org/xsd/maven-4.0.0.xsd">
	<modelVersion>4.0.0</modelVersion>
	<groupId>{{.Xml.GroupID}}</groupId>
	<artifactId>{{.Xml.ArtifactID}}</artifactId>
	<version>1.0</version>
	<name>{{.Name}}</name>
	<description>{{.Xml.Description}}</description>

	<repositories>
		<repository>
			<id>spigot-repo</id>
			<url>https://hub.spigotmc.org/nexus/content/repositories/snapshots/</url>
		</repository>
	</repositories>

	<properties>
		<maven.compiler.source>1.8</maven.compiler.source>
		<maven.compiler.target>1.8</maven.compiler.target>
	</properties>

	<dependencies>
		<dependency>
			<groupId>org.spigotmc</groupId>
			<artifactId>spigot-api</artifactId>
			<version>1.16.4-R0.1-SNAPSHOT</version>
			<scope>provided</scope>
		</dependency>
	</dependencies>

	<build>
		<sourceDirectory>src/main/java</sourceDirectory>
		<resources>
			<resource>
				<directory>src/main/java</directory>
				<includes>
					<include>plugin.yml</include>
				</includes>
			</resource>
		</resources>
		<plugins>
			<plugin>
				<groupId>org.apache.maven.plugins</groupId>
				<artifactId>maven-compiler-plugin</artifactId>
				<version>3.8.1</version>
			</plugin>
			<plugin>
				<artifactId>maven-assembly-plugin</artifactId>
				<configuration>
					<archive>
						<manifest>
							<mainClass>
								com.{{.Author}}.net.Main
							</mainClass>
						</manifest>
					</archive>
					<descriptorRefs>
						<descriptorRef>jar-with-dependencies</descriptorRef>
					</descriptorRefs>
				</configuration>
			</plugin>
		</plugins>
	</build>

</project>`
}