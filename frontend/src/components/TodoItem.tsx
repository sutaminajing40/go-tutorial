import React, { useState } from 'react'
import { Todo } from '../types/todo'

interface TodoItemProps {
  todo: Todo
  onToggle: (id: number) => void
  onEdit: (id: number, newTitle: string) => void
  onDelete: (id: number) => void
}

export default function TodoItem({ todo, onToggle, onEdit, onDelete }: TodoItemProps) {
  const [isEditing, setIsEditing] = useState(false)
  const [editedTitle, setEditedTitle] = useState(todo.title)

  const handleEdit = () => {
    setIsEditing(true)
  }

  const handleSave = () => {
    if (editedTitle.trim() !== '') {
      onEdit(todo.id, editedTitle.trim())
      setIsEditing(false)
    }
  }

  const handleCancel = () => {
    setEditedTitle(todo.title)
    setIsEditing(false)
  }

  return (
    <li className="flex items-center justify-between p-4 bg-white shadow rounded-lg mb-2">
      <div className="flex items-center">
        <input
          type="checkbox"
          checked={todo.completed}
          onChange={() => onToggle(todo.id)}
          className="h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
        />
        {isEditing ? (
          <input
            type="text"
            value={editedTitle}
            onChange={(e) => setEditedTitle(e.target.value)}
            className="ml-3 p-1 border rounded"
            autoFocus
          />
        ) : (
          <span className={`ml-3 ${todo.completed ? 'line-through text-gray-500' : 'text-gray-900'}`}>
            {todo.title}
          </span>
        )}
      </div>
      <div>
        {isEditing ? (
          <>
            <button
              onClick={handleSave}
              className="text-green-600 hover:text-green-800 mr-2"
            >
              Save
            </button>
            <button
              onClick={handleCancel}
              className="text-gray-600 hover:text-gray-800"
            >
              Cancel
            </button>
          </>
        ) : (
          <>
            <button
              onClick={handleEdit}
              className="text-blue-600 hover:text-blue-800 mr-2"
            >
              Edit
            </button>
            <button
              onClick={() => onDelete(todo.id)}
              className="text-red-600 hover:text-red-800"
            >
              Delete
            </button>
          </>
        )}
      </div>
    </li>
  )
}