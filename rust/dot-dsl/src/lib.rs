pub mod graph {
    pub use std::collections::HashMap;
    pub type Attrs = [(&'static str, &'static str)];

    pub mod graph_items {
        pub use super::{Attrs, HashMap};

        pub mod edge {
            use super::{Attrs, HashMap};

            #[derive(Debug, PartialEq, Eq, Clone)]
            pub struct Edge {
                pub source: String,
                pub target: String,
                pub attrs: HashMap<&'static str, &'static str>,
            }
            impl Edge {
                pub fn new(source: &str, target: &str) -> Self {
                    Edge {
                        source: source.to_string(),
                        target: target.to_string(),
                        attrs: HashMap::new(),
                    }
                }
                pub fn with_attrs(mut self, attrs: &Attrs) -> Self {
                    for &(key, val) in attrs.iter() {
                        self.attrs.insert(key, val);
                    }
                    self
                }
                pub fn attr(&self, name: &str) -> Option<&str> {
                    self.attrs.get(name).copied()
                }
            }
        }

        pub mod node {
            use super::{Attrs, HashMap};

            #[derive(Debug, PartialEq, Eq, Clone)]
            pub struct Node {
                pub id: &'static str,
                pub attrs: HashMap<&'static str, &'static str>,
            }

            impl Node {
                pub fn new(str: &'static str) -> Self {
                    Node {
                        id: str,
                        attrs: HashMap::new(),
                    }
                }
                pub fn with_attrs(mut self, attrs: &Attrs) -> Self {
                    self.attrs.extend(attrs.iter().cloned());
                    self
                }
                pub fn attr(&self, name: &'static str) -> Option<&str> {
                    self.attrs.get(name).copied()
                }
            }
        }
    }

    use graph_items::{edge::Edge, node::Node};
    #[derive(Debug)]
    pub struct Graph {
        pub nodes: Vec<Node>,
        pub edges: Vec<Edge>,
        pub attrs: HashMap<String, String>,
    }

    impl Graph {
        pub fn new() -> Self {
            Graph {
                nodes: Vec::new(),
                edges: Vec::new(),
                attrs: HashMap::new(),
            }
        }

        pub fn with_nodes(mut self, nodes: &[Node]) -> Self {
            self.nodes.extend(nodes.iter().cloned());
            self
        }
        pub fn with_edges(mut self, edges: &[Edge]) -> Self {
            self.edges.extend(edges.iter().cloned());
            self
        }
        pub fn with_attrs(mut self, attrs: &Attrs) -> Self {
            for (key, val) in attrs.iter() {
                self.attrs.insert(key.to_string(), val.to_string());
            }
            self
        }

        pub fn node(&self, id: &str) -> Option<&Node> {
            self.nodes.iter().find(|&node| node.id == id.to_string())
        }

        pub fn attr(&self, name: &str) -> Option<&str> {
            self.attrs.get(name).map(|x| x.as_str())
        }
    }
}
