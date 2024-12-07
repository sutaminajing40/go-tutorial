import React from 'react'

interface ErrorMessageProps {
  message: string
}

export default function ErrorMessage({ message }: ErrorMessageProps) {
  return (
    <div className="bg-red-100 border border-red-00 text-red-700 px-4 py-3 rounded relative mb-4" role="alert">
      <strong className="font-bold">Error: </strong>
      <span className="block sm:inline">{message}</span>
    </div>
  )
}
