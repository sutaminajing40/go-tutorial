import React, { useState, useEffect } from 'react'
import TodoList from './TodoList'
import TodoForm from './TodoForm'
import { Todo } from '../types/todo'
import { getTodos, createTodo, updateTodo, deleteTodo } from '../utils/api'
import LoadingSpinner from './LoadingSpinner'
import ErrorMessage from './ErrorMessage'

export default function TodoApp() {
  const [todos, setTodos] = useState<Todo[]>([])
  const [isLoading, setIsLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchTodos()
  }, [])

  const fetchTodos = async () => {
    try {
      setIsLoading(true)
      const fetchedTodos = await getTodos()
      setTodos(Array.isArray(fetchedTodos) ? fetchedTodos : [])
      setError(null)
    } catch (err) {
      setError(err as string)
      setTodos([])
    } finally {
      setIsLoading(false)
    }
  }

  const handleAddTodo = async (title: string) => {
    try {
      const newTodo = await createTodo(title)
      setTodos([...todos, newTodo])
    } catch (err) {
      setError(err as string)
    }
  }

  const handleToggleTodo = async (id: number) => {
    try {
      const todoToUpdate = todos.find(todo => todo.id === id)
      if (todoToUpdate) {
        const updatedTodo = await updateTodo(id, { completed: !todoToUpdate.completed })
        setTodos(todos.map(todo => todo.id === id ? updatedTodo : todo))
      }
    } catch (err) {
      setError(err as string)
    }
  }

  const handleEditTodo = async (id: number, newTitle: string) => {
    try {
      const updatedTodo = await updateTodo(id, { title: newTitle })
      setTodos(todos.map(todo => todo.id === id ? updatedTodo : todo))
    } catch (err) {
      setError(err as string)
    }
  }

  const handleDeleteTodo = async (id: number) => {
    try {
      await deleteTodo(id)
      setTodos(todos.filter(todo => todo.id !== id))
    } catch (err) {
      setError(err as string)
    }
  }

  if (isLoading) {
    return <LoadingSpinner />
  }

  return (
    <div>
      {error && <ErrorMessage message={error} />}
      <TodoForm onSubmit={handleAddTodo} />
      {Array.isArray(todos) ? (
        <TodoList
          todos={todos}
          onToggle={handleToggleTodo}
          onEdit={handleEditTodo}
          onDelete={handleDeleteTodo}
        />
      ) : (
        <ErrorMessage message="Unable to load todos. Please try again later." />
      )}
    </div>
  )
}

