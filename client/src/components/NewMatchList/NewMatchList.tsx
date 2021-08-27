import { VFC, useState, useEffect } from "react";
import {
  Flex,
  Stack,
  Text,
  Menu,
  MenuButton,
  MenuList,
  MenuItem,
  Button,
} from "@chakra-ui/react";
import { MAIN_COLOR } from "../../constants/MainColor";
import LoadingNewMatchList from "./LoadingNewMatchList";
import MatchListWrap from "../MatchListWrap/MatchListWrap";
import { useNewMatch } from "./useNewMatch";
import { ChevronDownIcon, CheckIcon } from "@chakra-ui/icons";
import {Match} from "../../type/Match"

const NewMatchList: VFC = () => {
  const { matches, error } = useNewMatch();
  const [matchList, setMatchList] = useState(matches);
  const [reactionType, setReactionType] = useState<"all" | "like" | "slike">("all");

  useEffect(() => {
    if (reactionType === "like") {
      setMatchList(matches.filter((match: Match) => match.reaction_type === 1));
    } else if (reactionType === "slike") {
      setMatchList(matches.filter((match: Match) => match.reaction_type === 2));
    } else setMatchList(matches);
  }, [reactionType, matches]);

  if (error) return <h1>An error has occurred.</h1>;
  if (!matchList) return <LoadingNewMatchList />;

  return (
    <Stack mb="4">
      <Flex align="center" justify="space-between">
        <Text
          fontSize="2xl"
          fontFamily={"heading"}
          color={`${MAIN_COLOR}.500`}
          fontWeight="bold"
          mr="4"
        >
          新しいマッチ
        </Text>
        <Menu>
          <MenuButton as={Button} rightIcon={<ChevronDownIcon />}>
            絞り込み
          </MenuButton>
          <MenuList>
            <MenuItem
              icon={
                <CheckIcon
                  color={reactionType === "all" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setReactionType("all")}
            >
              全ての企業
            </MenuItem>
            <MenuItem
              icon={
                <CheckIcon
                  color={reactionType === "like" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setReactionType("like")}
            >
              ライクした企業
            </MenuItem>
            <MenuItem
              icon={
                <CheckIcon
                  color={reactionType === "slike" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setReactionType("slike")}
            >
              スーパーライクした企業
            </MenuItem>
          </MenuList>
        </Menu>
      </Flex>
      <MatchListWrap matches={matches} />
    </Stack>
  );
};

export default NewMatchList;
