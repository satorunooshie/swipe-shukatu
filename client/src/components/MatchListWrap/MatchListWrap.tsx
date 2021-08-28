import { VFC } from "react";
import { Text, Image, VStack, Flex } from "@chakra-ui/react";
import { NavLink } from "react-router-dom";
import { Match } from "../../type/Match";
import { StarIcon } from "@chakra-ui/icons";

type Props = {
  readonly matches: Match[];
};

const MatchListWrap: VFC<Props> = ({ matches }) => {
  // const data: Match[] = [
  //   { ltd_id: 1, recruit_id: 1, name: "yahoo", image: "aaa", reaction_type: 1 },
  //   { ltd_id: 2, recruit_id: 2, name: "DeNa", image: "aaa", reaction_type: 2 },
  // ];
  return (
    <Flex wrap="nowrap" overflowX="auto">
      {matches.map((match: Match) => (
        <NavLink to={`/message/${match.recruit_id}`} key={match.recruit_id}>
          <VStack text="center" display="inline-block" mr="5" minW="100px">
            <Image
              w="100px"
              h="150px"
              borderRadius="lg"
              src={match.image}
              fit="cover"
            />
            <Flex justify="center" align="center">
              <Text align="center" fontWeight="bold">{match.name}</Text>
              {match.reaction_type === 2 && <StarIcon ml="1" color="cyan.400"/>}
            </Flex>
          </VStack>
        </NavLink>
      ))}
    </Flex>
  );
};

export default MatchListWrap;
