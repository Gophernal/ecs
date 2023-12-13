# ecs
An ECS implementation in Go for the Gophernal engine. Is generic enough to use for anything, though. 

# Installation

# Usage

# Implementation

## World

The world keeps track of the entities and systems, as well as match which entities go with which systems.
It does this using interfaces, so if an entity implements the right interface for a system (unless it
contains a blacklist flag!) then the system will act on it every update!

## Entities

Entities contain components and are acted on by systems. If the entity implements an interface the system
is looking for, then the system will act upon that entity every update.

## Components

Components are just the data. The data should be interacted with using methods, so that the component
can implement the proper interfaces when added to entities. For example, you should have something like
`SetPosition(x, y float32)` and `Move(dx, dy float32)` for a SpaceComponent rather than just a Position
property on the struct.

## Systems

Systems transform the data in the components when passed to them via their Update method.

## Benchmarks

