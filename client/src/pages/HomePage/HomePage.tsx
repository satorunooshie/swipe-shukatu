import "../../App.css";
import { VFC, useCallback, useContext } from "react";
import { Wrap, Box, Center, Flex, IconButton } from "@chakra-ui/react";
import Header from "../../components/Header/Header";
import CardContent from "../../components/CardContent/CardContent";
import Loading from "../../components/Loading/Loading";
import TinderCard from "react-tinder-card";
import { StarIcon, CloseIcon } from "@chakra-ui/icons";
import { FaHeart } from "react-icons/fa";
import { Ltd } from "../../type/Ltd";
import { useSWRInfinite } from "swr";
import { LoginModalContext } from "../../context/LoginModalContext";
import { CurrentUserContext } from "../../context/CurrentUserContext";
import axios from "../../lib/axios";
import { useShowToast } from "../../hooks/useShowToast";

// TODO: check
const fetcher = (url: string) =>
  axios
    .get(url, {
      headers: {
        Accept: "application/json",
      },
    })
    .then((res) => res.data);

const PAGE_SIZE = 20;

const HomePage: VFC = () => {
  const { onOpen } = useContext(LoginModalContext);
  const { currentUser } = useContext(CurrentUserContext);
  const showToast = useShowToast();
  // TODO: check
  const { data, error, mutate, size, setSize, isValidating } = useSWRInfinite(
    (index) => `/ltds?page=${index + 1}`,
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
    console.log(currentUser)
    // ユーザーはログインしていない
    if (currentUser === null || currentUser === undefined) {
      onOpen();
      showToast("ログインしてください", "error");
      return;
    }
    if (dir === "right") {
      like(ltd.id, currentUser.uid);
    } else if (dir === "left") {
      unlike(ltd.id, currentUser.uid);
    } else if (dir === "up") {
      superLike(ltd.id, currentUser.uid);
    }

    updateData(ltd.id);
  };
  // TODO: check
  const like = async (ltdId: string, currentUserId: string) => {
    axios
      .post("/like", {
        ltd_id: ltdId,
        headers: { Authorization: `Bearer ${currentUserId}` },
      })
      .then((res) => console.log(res))
      .catch((e) => console.log(e));
  };
  const unlike = async (ltdId: string, currentUserId: string) => {
    axios
      .post("/unlike", {
        ltd_id: ltdId,
        headers: { Authorization: `Bearer ${currentUserId}` },
      })
      .then((res) => console.log(res))
      .catch((e) => console.log(e));
  };
  const superLike = async (ltdId: string, currentUserId: string) => {
    axios
      .post("/superLike", {
        ltd_id: ltdId,
        headers: { Authorization: `Bearer ${currentUserId}` },
      })
      .then((res) => console.log(res))
      .catch((e) => console.log(e));
  };

  if (error)
    return (
      <Center h="full" w="full">
        <Wrap h="90vh" w="full">
          <Header />
          <Box w="full" h="80vh">
            <Center m="auto" w="90vh" maxW="300px" h="400px">
              Error
            </Center>
          </Box>
        </Wrap>
      </Center>
    );
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
        <Box w="full" position="relative" h="80vh" pt={["10%", "5%", 0]}>
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
          <Flex
            w={["45%", "45%", "35%"]}
            justify="space-between"
            align="center"
          >
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
