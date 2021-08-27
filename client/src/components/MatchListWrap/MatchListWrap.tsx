import { VFC } from "react";
import { Text, Image, VStack, Flex } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";
import { Ltd } from "../../type/Ltd";

type Props = {
  readonly ltds: Ltd[];
};

const MatchListWrap: VFC<Props> = ({ ltds }) => {
  return (
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
  );
};

export default MatchListWrap;
