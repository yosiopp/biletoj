import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import './index.css'
import TicketList from "./pages/TicketList";
import TicketDetail from "./pages/TicketDetail";
import { useEffect } from "react";

const router = createBrowserRouter([
  {
    path: "/tickets",
    element: <TicketList />,
  },
  {
    path: "/tickets/:id",
    element: <TicketDetail />,
  },
  {
    path: "/*",
    element: <TicketList />,
  },
]);

function App() {
  useEffect(() => {
    const eventListener = (event: KeyboardEvent) => {
      // TODO
      if (event.ctrlKey) {
        console.log(event);
        event.preventDefault();
      }
    };
    window.addEventListener("keydown", eventListener);
    return () => window.removeEventListener("keydown", eventListener);
  })
  return (
    <RouterProvider router={router} />
  )
}

export default App
