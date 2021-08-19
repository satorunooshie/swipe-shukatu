import { VFC } from "react";
import { Flex, Stack, Text, Container, Box, VStack } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";

const NewMatchList: VFC = () => {
  return (
    <Stack>
      <Text
        fontSize="2xl"
        fontFamily={"heading"}
        color={`${MAIN_COLOR}.500`}
        fontWeight="bold"
        mb="4"
      >
        新しいマッチ
      </Text>
      <Flex wrap="nowrap" overflowX="auto">
        {["A社", "B社", "C社", "D社", "E社", "F社", "G社"].map((ltd, i) => (
          <VStack text="center" display="inline-block" mr="5">
            <Box bg="gray.400" w="100px" h="150px" borderRadius="lg"></Box>
            <Text align="center">{ltd}</Text>
          </VStack>
        ))}
      </Flex>
    </Stack>
  );
};

export default NewMatchList;
