import React from 'react'
import TodoItem from './TodoItem'
import { Todo } from '../types/todo'

interface TodoListProps {
  todos: Todo[]
  onToggle: (id: number) => void
  onEdit: (id: number) => void
  onDelete: (id: number) => void
}

export default function TodoList({ todos, onToggle, onEdit, onDelete }: TodoListProps) {
  return (
    <ul className="space-y-2">
      {todos.map((todo) => (
        <TodoItem
          key={todo.id}
          todo={todo}
          onToggle={onToggle}
          onEdit={onEdit}
          onDelete={onDelete}
        />
      ))}
    </ul>
  )
}

