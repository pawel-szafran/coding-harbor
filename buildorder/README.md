# Find build order in a multi-module project

Build dependencies first. And detect cycles!

### Example

Modules with dependencies:
```
A -> {B, C, F}
B -> {F}
C -> {F}
D -> {}
E -> {A, B}
F -> {}
G -> {D}
```
can be built as: `[F, B, C, A, D, E, G]`

### Solution

Topological sort with recursive DFS.
