import React from "react";
import "./App.css";
import { ChakraProvider, Text, Divider } from "@chakra-ui/react";
import { BrowserRouter as Router } from "react-router-dom";
import Navigation from "../src/components/Navigation/Navigation";

function App() {
  return (
    <ChakraProvider>
      <Navigation />
      <Router>
        <Text>Edit</Text>
      </Router>
    </ChakraProvider>
  );
}

export default App;
