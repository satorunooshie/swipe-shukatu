import { VFC } from "react";
import {
  Flex,
  Divider,
  Wrap,
  Spacer,
  Text,
  Button,
  IconButton,
} from "@chakra-ui/react";
import { NavLink, useLocation } from "react-router-dom";
import { SearchIcon } from "@chakra-ui/icons";

const Header: VFC = () => {
  let location = useLocation();
  if (location.pathname !== "/") return <></>;
  return (
    <Wrap pos="sticky" w="full">
      <Flex
        w="full"
        minH={"50px"}
        mb="60px"
        py={{ base: 2 }}
        px={{ base: 10 }}
        borderBottom={1}
        align={"center"}
        fontWeight="bold"
      >
        <NavLink to="/">
          <Text fontFamily={"heading"} color={"gray.800"}>
            Swipe Shukatsu
          </Text>
        </NavLink>
        <Spacer />
        <IconButton
          aria-label="Search"
          variant="goast"
          mr="4"
          icon={<SearchIcon />}
        />
        <Button
          fontSize={"sm"}
          fontWeight={600}
          color={"white"}
          bg={"blue.400"}
          _hover={{
            bg: "blue.300",
          }}
        >
          Log In
        </Button>
      </Flex>
      <Divider />
    </Wrap>
  );
};

export default Header;
