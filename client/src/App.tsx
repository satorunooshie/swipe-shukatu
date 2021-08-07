import React from "react";
import "./App.css";
import { ChakraProvider, Text } from "@chakra-ui/react";
import { BrowserRouter as Router } from "react-router-dom";

function App() {
  return (
    <ChakraProvider>
      <Router>
        <Text>Edit</Text>
      </Router>
    </ChakraProvider>
  );
}

export default App;
