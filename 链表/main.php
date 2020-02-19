<?php

class Node
{
    public $val;
    public $next;
}

// 翻转链表
function reverse($node)
{
    $old = $node->next;
    $new = '';
    while (!empty($old)) {
        $old_tmp = $old->next;
        $old->next = $new;
        $new = $old;
        $old = $old_tmp;
    }
    return $new;
}

// 创建1-10的链表
$node = new Node();
for ($i = 1; $i <= 10; $i++) {
    $node_list = new Node();
    $node_list->val = $i;
    $node_list->next = $node->next;
    $node->next = $node_list;
}
// print_r($node);

print_r(reverse($node));
