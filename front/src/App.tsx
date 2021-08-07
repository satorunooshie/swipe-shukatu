import React from "react";
import "./App.css";
import { ChakraProvider, Text } from "@chakra-ui/react";

function App() {
  return (
    <ChakraProvider>
      <Text>Edit</Text>
    </ChakraProvider>
  );
}

export default App;
