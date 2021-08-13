// @ts-nocheck
import "../../App.css";
import { VFC, useMemo, createRef, useState, useEffect } from "react";
import { Wrap, Box, Center, Flex, IconButton } from "@chakra-ui/react";
import Header from "../../components/Header/Header";
import TinderCard from "react-tinder-card";
import { StarIcon, CloseIcon } from "@chakra-ui/icons";
import { FaHeart } from "react-icons/fa";

// 仮
type Ltd = {
  readonly id: number;
  readonly name: string;
};

const initialLtds: Ltd[] = [
  {
    id: 1,
    name: "会社A",
  },
  {
    id: 2,
    name: "会社B",
  },
  {
    id: 3,
    name: "会社C",
  },
  {
    id: 4,
    name: "会社D",
  },
  {
    id: 5,
    name: "会社E",
  },
  {
    id: 6,
    name: "会社F",
  },
  {
    id: 7,
    name: "会社G",
  },
  {
    id: 8,
    name: "会社H",
  },
  {
    id: 9,
    name: "会社I",
  },
  {
    id: 10,
    name: "会社J",
  },
];

const ltds2: Ltd[] = [
  {
    id: 11,
    name: "会社K",
  },
  {
    id: 12,
    name: "会社L",
  },
  {
    id: 13,
    name: "会社M",
  },
  {
    id: 14,
    name: "会社N",
  },
  {
    id: 15,
    name: "会社O",
  },
];

const ltds3: Ltd[] = [
  {
    id: 16,
    name: "会社P",
  },
  {
    id: 17,
    name: "会社Q",
  },
  {
    id: 18,
    name: "会社R",
  },
  {
    id: 19,
    name: "会社S",
  },
  {
    id: 20,
    name: "会社T",
  },
];
const Ltds: Ltd[] = [ltds2, ltds3];

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
    console.log(cardsLeft);
    console.log(childRefs);
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
    console.log(childRefs);
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
      const toBeRemoved = cardsLeft[cardsLeft.length - 1].id; // Find the card object to be removed
      const index = ltdList.map((ltd) => ltd.id).indexOf(toBeRemoved); // Find the index of which to make the reference to

      setAlreadyRemoved([...alreadyRemoved, toBeRemoved]);
      childRefs[index].current.swipe(dir); // Swipe the card!

      // console.log("cardsLeft", cardsLeft);
      // console.log("toBeRemoved", toBeRemoved);
      // console.log("index", index);
      // console.log("childRrefs", childRefs);
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
                key={ltd.name}
                className="swipe"
                ref={childRefs[index]}
                onSwipe={(dir) => swiped(dir, ltd)}
                preventSwipe={["down"]}
              >
                <Box
                  className="card"
                  border="1px"
                  borderColor="gray.300"
                  pos="relative"
                  bg="gray.100"
                  h="500"
                  w="80vh"
                  maxW="300"
                  borderRadius="10"
                >
                  <h3>{ltd.name}</h3>
                </Box>
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
