import React from "react";
import "./App.css";
import { ChakraProvider, VStack, Flex } from "@chakra-ui/react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import BottomTabs from "./components/BottomTabs/BottomTabs";
import Sidebar from "../src/components/Sidebar/Sidebar";
import LoginModal from "../src/components/LoginModal/LoginModal";
import HomePage from "../src/pages/HomePage/HomePage";
import MessagePage from "../src/pages/MessagePage/MessagePage";
import ChatPage from "../src/pages/ChatPage/ChatPage";
import RecommendPage from "../src/pages/RecommendPage/RecommendPage";
import { LoginModalProvider } from "../src/context/LoginModalContext";
import { CurrentUserProvider } from "./context/CurrentUserContext";

function App() {
  return (
    <ChakraProvider>
      <CurrentUserProvider>
        <LoginModalProvider>
          <Router>
            <LoginModal />
            <VStack px={[0, 0, 0, "10%"]}>
              <BottomTabs />
              <Flex w="full" justify="center">
                <Sidebar />
                <Switch>
                  <Route path="/recommend">
                    <RecommendPage />
                  </Route>
                  <Route path="/message/:recruit_id">
                    <ChatPage />
                  </Route>
                  <Route path="/message">
                    <MessagePage />
                  </Route>
                  <Route path="/">
                    <HomePage />
                  </Route>
                </Switch>
              </Flex>
            </VStack>
          </Router>
        </LoginModalProvider>
      </CurrentUserProvider>
    </ChakraProvider>
  );
}

export default App;
