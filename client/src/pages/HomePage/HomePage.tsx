import "../../App.css";
import { VFC, useCallback } from "react";
import {
  Wrap,
  Box,
  Center,
  Flex,
  IconButton,
} from "@chakra-ui/react";
import Header from "../../components/Header/Header";
import CardContent from "../../components/CardContent/CardContent";
import Loading from "../../components/Loading/Loading";
import TinderCard from "react-tinder-card";
import { StarIcon, CloseIcon } from "@chakra-ui/icons";
import { FaHeart } from "react-icons/fa";
import { Ltd } from "../../type/Ltd";
import { useSWRInfinite } from "swr";

const fetcher = (url: string) =>
  fetch(url, {
    headers: {
      Accept: "application/json",
    },
  })
    .then((res) => res.json())
    .then((res) => res.results);

const PAGE_SIZE = 20;

const HomePage: VFC = () => {
  const { data, error, mutate, size, setSize, isValidating } = useSWRInfinite(
    (index) => `https://icanhazdadjoke.com/search?page=${index + 1}`,
    fetcher
  );
  const alreadyRemoved = new Set<string>();

  const updateData = useCallback(async (ltdID: string) => {
    alreadyRemoved.add(ltdID);

    // 残りが5件以下になったらデータフェッチする
    if (PAGE_SIZE - alreadyRemoved.size > 5) return;
    alreadyRemoved.clear();
    setSize((size) => size + 1);
  }, []);

  const swiped = (dir: "right" | "left" | "up" | "down", ltd: Ltd) => {
    if (dir === "right") {
      console.log("like", ltd.joke);
    } else if (dir === "left") {
      console.log("nope", ltd.joke);
    } else if (dir === "up") {
      console.log("slike", ltd.joke);
    }

    updateData(ltd.id);
  };

  if (error) return <h1>error occured</h1>;
  if (!data)
    return (
      <Center h="full" w="full">
        <Wrap h="90vh" w="full">
          <Header />
          <Box w="full" h="80vh">
            <Center m="auto" w="90vh" maxW="300px" h="400px">
              <Loading />
            </Center>
          </Box>
        </Wrap>
      </Center>
    );

  return (
    <Center h="full" w="full">
      <Wrap h="90vh" w="full">
        <Header />
        <Box w="full" position="relative" h="80vh">
          <Wrap m="auto" w="90vh" maxW="300px" h="300px">
            {data.map((ltds) =>
              ltds.map((ltd: Ltd, index: number) => (
                <TinderCard
                  key={index}
                  // @ts-ignore
                  className="swipe"
                  // @ts-ignore
                  // ref={refs[index]}
                  onSwipe={(dir) => swiped(dir, ltd)}
                  preventSwipe={["down"]}
                >
                  <CardContent ltd={ltd} />
                </TinderCard>
              ))
            )}
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
              // onClick={() => swipe("left")}
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
              // onClick={() => swipe("up")}
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
              // onClick={() => swipe("right")}
            />
          </Flex>
        </Wrap>
      </Wrap>
    </Center>
  );
};

export default HomePage;
