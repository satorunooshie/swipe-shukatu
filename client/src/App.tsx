import React from "react";
import "./App.css";
import { ChakraProvider, Text } from "@chakra-ui/react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Navigation from "../src/components/Navigation/Navigation";
import HomePage from "../src/pages/HomePage/HomePage";
import MessagePage from "../src/pages/MessagePage/MessagePage";
import RecommendPage from "../src/pages/RecommendPage/RecommendPage";

function App() {
  return (
    <ChakraProvider>
      <Router>
        <Navigation />
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
      </Router>
    </ChakraProvider>
  );
}

export default App;
