package model

type Node struct {
	Key    int   `json:"key"`
	Left   *Node `json:"left,omitempty"`
	Right  *Node `json:"right,omitempty"`
	Parent *Node `json:"-"`
}
