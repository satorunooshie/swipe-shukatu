import { VFC } from "react";
import { Flex, Stack, Text, Image, VStack } from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import { NavLink } from "react-router-dom";
import { Ltd } from "../../type/Ltd";
import LoadingNewMatchList from "./LoadingNewMatchList";
import { useNewMatch } from "./useNewMatch";

const NewMatchList: VFC = () => {
  const { ltds, error } = useNewMatch();

  if (error) return <h1>An error has occurred.</h1>;
  if (!ltds) return <LoadingNewMatchList />;

  return (
    <Stack mb="4">
      <Text
        fontSize="2xl"
        fontFamily={"heading"}
        color={`${MAIN_COLOR}.500`}
        fontWeight="bold"
      >
        新しいマッチ
      </Text>
      <Flex wrap="nowrap" overflowX="auto">
        {ltds.map((ltd: Ltd) => (
          <NavLink to={`/message/${ltd.id}`} key={ltd.id}>
            <VStack text="center" display="inline-block" mr="5" minW="100px">
              <Image
                w="100px"
                h="150px"
                borderRadius="lg"
                src={`https://icanhazdadjoke.com/j/${ltd.id}.png`}
                fit="cover"
              />
              <Text align="center">{ltd.joke.slice(0, 10)}</Text>
            </VStack>
          </NavLink>
        ))}
      </Flex>
    </Stack>
  );
};

export default NewMatchList;
