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
import { Ltd } from "../../type/Ltd";
import LoadingNewMatchList from "./LoadingNewMatchList";
import MatchListWrap from "../MatchListWrap/MatchListWrap";
import { useNewMatch } from "./useNewMatch";
import { ChevronDownIcon, CheckIcon } from "@chakra-ui/icons";

const NewMatchList: VFC = () => {
  const { ltds, error } = useNewMatch();
  const [ltdList, setLtdList] = useState(ltds);
  const [ltdType, setLtdType] = useState<"all" | "like" | "slike">("all");

  useEffect(() => {
    if (ltdType === "like") {
      // Todo:
      // setLtdList(ltds.filter((ltd: Match) => ltd.reaction_type === 1));
      setLtdList(ltds.filter((ltd: Ltd) => ltd.joke.length % 2 === 0));
    } else if (ltdType === "slike") {
      setLtdList(ltds.filter((ltd: Ltd) => ltd.joke.length % 2 === 1));
    } else setLtdList(ltds);
  }, [ltdType, ltds]);

  if (error) return <h1>An error has occurred.</h1>;
  if (!ltdList) return <LoadingNewMatchList />;

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
                  color={ltdType === "all" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setLtdType("all")}
            >
              全ての企業
            </MenuItem>
            <MenuItem
              icon={
                <CheckIcon
                  color={ltdType === "like" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setLtdType("like")}
            >
              ライクした企業
            </MenuItem>
            <MenuItem
              icon={
                <CheckIcon
                  color={ltdType === "slike" ? "gray.600" : "gray.100"}
                />
              }
              onClick={() => setLtdType("slike")}
            >
              スーパーライクした企業
            </MenuItem>
          </MenuList>
        </Menu>
      </Flex>
      <MatchListWrap ltds={ltdList} />
    </Stack>
  );
};

export default NewMatchList;
