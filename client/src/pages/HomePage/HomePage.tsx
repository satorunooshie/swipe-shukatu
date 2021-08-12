// @ts-nocheck
import "../../App.css";
import { VFC } from "react";
import {
  Button,
  Wrap,
  Box,
  Container,
  Center,
  Flex,
  IconButton,
} from "@chakra-ui/react";
import Header from "../../components/Header/Header";
import TinderCard from "react-tinder-card";
import { StarIcon,CloseIcon } from "@chakra-ui/icons";
import {FaHeart} from "react-icons/fa"
const ltds = [
  {
    name: "会社A",
  },
  {
    name: "会社B",
  },
  {
    name: "会社C",
  },
  {
    name: "会社D",
  },
  {
    name: "会社E",
  },
];

const HomePage: VFC = () => {
  return (
    <Center h="full" w="full">
      <Wrap h="90vh" w="full">
        <Header />
        <Box w="full" position="relative" h="80vh">
          <Wrap m="auto" w="90vh" maxW="300px" h="300px">
            {ltds.map((character, index) => (
              <TinderCard key={character.name} className="swipe">
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
                  <h3>{character.name}</h3>
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
              icon={<CloseIcon w={5} h={5}/>}
            />
            <IconButton
              size="md"
              variant="outline"
              isRound
              bg="transparent"
              color="cyan.300"
              colorScheme="cyan"
              aria-label="super like"
              icon={<StarIcon w={5} h={5}/>}
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
            />
          </Flex>
        </Wrap>
      </Wrap>
    </Center>
  );
};

export default HomePage;
