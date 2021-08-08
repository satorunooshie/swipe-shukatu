import { VFC } from "react";
import {
  Flex,
  IconButton,
  Center,
  Icon,
  Divider,
  Wrap,
  VStack,
} from "@chakra-ui/react";
import { AiFillFire } from "react-icons/ai";
import { HiChatAlt2 } from "react-icons/hi";
import { BsLightningFill } from "react-icons/bs";
import { NavLink, useLocation } from "react-router-dom";

const Sidebar: VFC = () => {
  let location = useLocation();

  return (
      <VStack
      display={["none", "none", "block"]}
      pos="static"
      h="75vh"
      w="200px"
      marginTop="2.5vh"
      boxShadow="0 4px 12px 0 rgba(0, 0, 0, 0.05)"
      borderRadius="15px"
      bg="red"
      >

      </VStack>
  );
};

export default Sidebar;
