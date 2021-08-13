import { VFC } from "react";
import {
  Wrap,
  Box,
  Flex,
  Spacer,
  Heading,
  Tag,
  useDisclosure,
  Modal,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  Stack,
  StackDivider,
  Text,
  Center,
} from "@chakra-ui/react";
import { Ltd } from "../../type/Ltd";
import { MAIN_COLOR } from "../../constants/MainColor";

type Props = {
  readonly ltd: Ltd;
};

const CardContent: VFC<Props> = ({ ltd }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <Modal isOpen={isOpen} onClose={onClose} scrollBehavior="outside">
        <ModalContent>
          <ModalHeader>
            <Flex alignItems="center">
              <Heading color="gray.600">{ltd.name}</Heading>
              <Text color="gray.500" fontSize="3xl" ml="4">
                IT
              </Text>
            </Flex>
          </ModalHeader>
          <ModalBody>
            <Stack
              divider={<StackDivider borderColor="gray.200" />}
            >
              <StackDivider borderColor="gray.200" />
              <Wrap my="2">
                <Text color="gray.500">830万円</Text>
              </Wrap>
              <Wrap my="5">
                <Text color="gray.500">
                  (会社紹介)Lorem ipsum is placeholder text commonly used in the
                  graphic, print, and publishing industries for previewing
                  layouts and visual mockups.
                </Text>
              </Wrap>
              <Wrap my="5">
                <Text color="gray.700">会社情報</Text>
                <Wrap w="full">
                  <Tag colorScheme={MAIN_COLOR}>家賃補助</Tag>
                  <Tag colorScheme={MAIN_COLOR}>技術書購入手当</Tag>
                  <Tag colorScheme={MAIN_COLOR}>海外カンファレンス補助</Tag>
                  <Tag colorScheme={MAIN_COLOR}>お子さんのバースデー休暇</Tag>
                </Wrap>
              </Wrap>
              <Wrap textAlign="center" w="full">
                <Center w="full">
                  <Text color="gray.500">{ltd.name}を報告する</Text>
                </Center>
              </Wrap>
            </Stack>
          </ModalBody>
          <ModalFooter></ModalFooter>
        </ModalContent>
      </Modal>
      <Box
        className="card"
        border="1px"
        borderColor="gray.300"
        pos="relative"
        bg="gray.800"
        h="500"
        w="80vh"
        maxW="300"
        borderRadius="10"
      >
        <Flex flexDirection="column" w="full" h="full">
          <Spacer />
          <Wrap w="full" h={32} p={3}>
            <Heading
              as="h2"
              size="lg"
              color="white"
              textShadow="1px 0px 5px #000"
              onClick={() => onOpen()}
            >
              {ltd.name}
            </Heading>
            <Wrap w="full">
              <Tag colorScheme={MAIN_COLOR}>Tag</Tag>
              <Tag colorScheme={MAIN_COLOR}>Tag</Tag>
              <Tag colorScheme={MAIN_COLOR}>Tag</Tag>
            </Wrap>
          </Wrap>
        </Flex>
      </Box>
    </>
  );
};

export default CardContent;
