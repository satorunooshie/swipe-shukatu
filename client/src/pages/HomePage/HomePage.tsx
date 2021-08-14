import "../../App.css";
import { VFC, createRef, useState, useEffect } from "react";
import { Wrap, Box, Center, Flex, IconButton } from "@chakra-ui/react";
import Header from "../../components/Header/Header";
import CardContent from "../../components/CardContent/CardContent";
import TinderCard from "react-tinder-card";
import { StarIcon, CloseIcon } from "@chakra-ui/icons";
import { FaHeart } from "react-icons/fa";
import { Ltd } from "../../type/Ltd";
import { initialLtds, Ltds} from "./DB"

const HomePage: VFC = () => {
  const [ltdList, setLtdList] = useState<Ltd[]>(initialLtds);
  const [cardsLeft, setCardsLeft] = useState<Ltd[]>(initialLtds);
  const [alreadyRemoved, setAlreadyRemoved] = useState<number[]>([]);
  const [childRefs, setChildRefs] = useState(
    Array(ltdList.length)
      .fill(0)
      .map((i) => createRef())
  );
  const [page, setPage] = useState(0);

  useEffect(() => {
    // cardsLeftを算出
    setCardsLeft(ltdList.filter((ltd) => !alreadyRemoved.includes(ltd.id)));
  }, [alreadyRemoved]);

  useEffect(() => {
    if (cardsLeft.length > 5) return;

    // TODO: ここでデータフェッチ,ここで言う newDataに割り当てる
    const newData = Ltds[page];
    if (page >= 2) return;
    setPage(page + 1);
    setLtdList([...ltdList, ...newData]);
    setCardsLeft([...cardsLeft, ...newData]);
    setChildRefs([
      ...childRefs,
      ...Array(newData.length)
        .fill(0)
        .map((i) => createRef()),
    ]);
  }, [cardsLeft]);

  const swiped = (dir: "right" | "left" | "up" | "down", ltd: Ltd) => {
    if (dir === "right") {
      console.log("like", ltd.name);
    } else if (dir === "left") {
      console.log("nope", ltd.name);
    } else if (dir === "up") {
      console.log("slike", ltd.name);
    }
    setAlreadyRemoved([...alreadyRemoved, ltd.id]);
  };

  const swipe = (dir: "right" | "left" | "up" | "down") => {
    if (cardsLeft.length) {
      const toBeRemoved = cardsLeft[cardsLeft.length - 1].id;
      const index = ltdList.map((ltd) => ltd.id).indexOf(toBeRemoved);

      setAlreadyRemoved([...alreadyRemoved, toBeRemoved]);
      // @ts-ignore
      childRefs[index].current.swipe(dir);
    }
  };

  if (cardsLeft.length === 0) return <h1>All swiped...</h1>;

  return (
    <Center h="full" w="full">
      <Wrap h="90vh" w="full">
        <Header />
        <Box w="full" position="relative" h="80vh">
          <Wrap m="auto" w="90vh" maxW="300px" h="300px">
            {ltdList.map((ltd, index) => (
              <TinderCard
                key={ltd.id}
                className="swipe"
                // @ts-ignore
                ref={childRefs[index]}
                onSwipe={(dir) => swiped(dir, ltd)}
                preventSwipe={["down"]}
              >
                <CardContent ltd={ltd} />
              </TinderCard>
            ))}
          </Wrap>
        </Box>
        <Wrap w="full" pos="sticky" bottom="85px" justify="center">
          <Flex w="35%" justify="space-between" align="center">
            <IconButton
              size="lg"
              variant="outline"
              isRound
              bg="transparent"
              color="pink.300"
              colorScheme="pink"
              aria-label="super like"
              icon={<CloseIcon w={5} h={5} />}
              onClick={() => swipe("left")}
            />
            <IconButton
              size="md"
              variant="outline"
              isRound
              bg="transparent"
              color="cyan.300"
              colorScheme="cyan"
              aria-label="super like"
              icon={<StarIcon w={5} h={5} />}
              onClick={() => swipe("up")}
            />
            <IconButton
              size="lg"
              variant="outline"
              isRound
              bg="transparent"
              color="teal.300"
              colorScheme="teal"
              aria-label="super like"
              icon={<FaHeart />}
              onClick={() => swipe("right")}
            />
          </Flex>
        </Wrap>
      </Wrap>
    </Center>
  );
};

export default HomePage;
