import React from "react";
import "./App.css";
import { ChakraProvider, Center, Stack, Flex } from "@chakra-ui/react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Navigation from "../src/components/Navigation/Navigation";
import Sidebar from "../src/components/Sidebar/Sidebar";
import Header from "../src/components/Header/Header";
import HomePage from "../src/pages/HomePage/HomePage";
import MessagePage from "../src/pages/MessagePage/MessagePage";
import RecommendPage from "../src/pages/RecommendPage/RecommendPage";

function App() {
  return (
    <ChakraProvider>
      <Router>
        <Center px={[0, 0, 0, "10%"]} >
          <Stack w="full">
            <Header />
            <Navigation />
            <Flex w="full" justify="center">
              <Sidebar />
              <Switch>
                <Route path="/recommend">
                  <RecommendPage />
                </Route>
                <Route path="/message">
                  <MessagePage />
                </Route>
                <Route path="/">
                  <HomePage />
                </Route>
              </Switch>
            </Flex>
          </Stack>
        </Center>
      </Router>
    </ChakraProvider>
  );
}

export default App;
