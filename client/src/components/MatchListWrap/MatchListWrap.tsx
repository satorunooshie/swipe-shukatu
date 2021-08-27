import { VFC } from "react";
import { Text, Image, VStack, Flex } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";
import { Match } from "../../type/Match";

type Props = {
  readonly matches: Match[];
};

const MatchListWrap: VFC<Props> = ({ matches }) => {
  return (
    <Flex wrap="nowrap" overflowX="auto">
      {matches.map((match: Match) => (
        <NavLink to={`/message/${match.ltd_id}`} key={match.ltd_id}>
          <VStack text="center" display="inline-block" mr="5" minW="100px">
            <Image
              w="100px"
              h="150px"
              borderRadius="lg"
              src={match.image}
              fit="cover"
            />
            <Text align="center">{match.name}</Text>
          </VStack>
        </NavLink>
      ))}
    </Flex>
  );
};

export default MatchListWrap;
