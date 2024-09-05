import React from 'react';
import './index.css';
import * as ReactDOM from "react-dom/client"
import {
  createBrowserRouter,
  RouterProvider  
} from "react-router-dom"
import PostEditor from './postEditor/postEditor';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

const router = createBrowserRouter([
  {
    path: "/",
    element: <div>Hello world!</div>
  },
  {
    path: "/post-editor",
    element: <PostEditor />
  }
])

root.render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>
);