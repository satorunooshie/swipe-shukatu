import { VFC } from "react";
import {
  Wrap,
  Box,
  Flex,
  Spacer,
  Heading,
  Tag,
  useDisclosure,
  Icon,
} from "@chakra-ui/react";
import { Ltd } from "../../type/Ltd";
import { MAIN_COLOR } from "../../constants/MainColor";
import LtdDetailModal from "../LtdDetailModal/LtdDetailModal";

type Props = {
  readonly ltd: Ltd;
};

const CardContent: VFC<Props> = ({ ltd }) => {
  const { isOpen, onOpen, onClose } = useDisclosure();

  return (
    <>
      <LtdDetailModal ltd={ltd} isOpen={isOpen} onClose={onClose} />
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
          <Wrap w="full" h={32} p={3} onClick={() => onOpen()}>
            <Heading
              as="h2"
              size="lg"
              color="white"
              textShadow="1px 0px 5px #000"
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
