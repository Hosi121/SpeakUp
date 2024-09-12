import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import App from "./App.tsx";
import { Settings } from "./components/pages/Settings.tsx";
import { Home } from "./components/pages/Home.tsx";
import { SessionList } from "./components/pages/SessionList.tsx";
import { Waiting } from "./components/pages/Waiting.tsx";
import { MicCheck } from "./components/pages/MicCheck.tsx";
import { Session } from "./components/pages/Session.tsx";
import { SessionInterval } from "./components/pages/SessionInterval.tsx";
import { SessionFinished } from "./components/pages/SessionFinished.tsx";
import { SignUp } from "./components/pages/SignUp.tsx";
import { Record } from "./components/pages/Record.tsx";
import { Memo } from "./components/pages/Memo.tsx";
import { Message } from "./components/pages/Message.tsx";
import { Login } from "./components/pages/Login.tsx";
import { Alert } from "./components/pages/Alert.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "alert",
    element: <Alert />,
  },
  {
    path: "settings",
    element: <Settings />,
  },
  {
    path: "login",
    element: <Login />,
  },
  {
    path: "home",
    element: <Home />,
  },
  {
    path: "sessionlist",
    element: <SessionList />,
  },
  {
    path: "waiting",
    element: <Waiting />,
  },
  {
    path: "miccheck",
    element: <MicCheck />,
  },
  {
    path: "session",
    element: <Session />,
  },
  {
    path: "sessioninterval",
    element: <SessionInterval />,
  },
  {
    path: "sessionfinished",
    element: <SessionFinished />,
  },
  {
    path: "signup",
    element: <SignUp />,
  },
  {
    path: "record",
    element: <Record />,
  },
  {
    path: "memo",
    element: <Memo />,
  },
  {
    path: "message",
    element: <Message />,
  },
]);

const rootElement = document.getElementById("root");
if (rootElement) {
  createRoot(rootElement).render(
    <StrictMode>
      <RouterProvider router={router} />
    </StrictMode>
  );
} else {
  console.error("Failed to find the root element");
}
