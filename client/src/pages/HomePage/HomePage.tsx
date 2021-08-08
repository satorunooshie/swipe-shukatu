import { VFC } from "react";
import { Text, Wrap } from "@chakra-ui/react";
import Header from "../../components/Header/Header";

const HomePage: VFC = () => {
  return (
    <Wrap h="90vh" w="full">
      <Header />
      <Text>Home</Text>
    </Wrap>
  );
};

export default HomePage;
