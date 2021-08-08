import { VFC } from "react";
import {
  Flex,
  Spacer,
  Text,
  Button,
} from "@chakra-ui/react";
import { NavLink, useLocation } from "react-router-dom";
import { SearchIcon } from "@chakra-ui/icons";
import { MAIN_COLOR } from "../../constants/MainColor";

const Header: VFC = () => {
  let location = useLocation();
  if (location.pathname !== "/") return <></>;
  return (
    <Flex
      pos="sticky"
      w="full"
      minH={"50px"}
      mb="60px"
      py={{ base: 2 }}
      px={{ base: 4 }}
      borderBottom={1}
      align={"center"}
      justify="center"
    >
      <NavLink to="/">
        <Text
          display={["block", "block", "none"]}
          fontFamily={"heading"}
          color={"gray.800"}
          fontWeight="bold"
        >
          Swipe Shukatsu
        </Text>
      </NavLink>
      <Spacer display={{ md: "none" }} />
      <Button
        onClick={() => alert("ここクリックするとモーダル開く")}
        bg="gray.100"
        px="6"
        py="2"
        borderRadius="lg"
      >
        <Flex align="center" w={{ md: "sm" }}>
          <SearchIcon />
          <Text ml="4">検索条件</Text>
        </Flex>
      </Button>
      <Button
        display={["block", "block", "none"]}
        fontSize={"sm"}
        fontWeight={600}
        color={"white"}
        ml="4"
        bg={`${MAIN_COLOR}.400`}
        _hover={{
          bg: `${MAIN_COLOR}.300`,
        }}
      >
        Log In
      </Button>
    </Flex>
  );
};

export default Header;
