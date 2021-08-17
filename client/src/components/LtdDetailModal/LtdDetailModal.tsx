import { VFC } from "react";
import {
  Wrap,
  Flex,
  Heading,
  Tag,
  Modal,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  Stack,
  StackDivider,
  Text,
  Center,
  Icon,
  ModalCloseButton,
  Divider,
} from "@chakra-ui/react";
import { Ltd } from "../../type/Ltd";
import { FcAdvertising, FcMoneyTransfer, FcInfo } from "react-icons/fc";

type Props = {
  readonly ltd: Ltd;
  readonly isOpen: boolean;
  readonly onClose: () => void;
};

const LtdDetailModal: VFC<Props> = ({ ltd, isOpen, onClose }) => {
  return (
    <Modal isOpen={isOpen} onClose={onClose} scrollBehavior="outside">
      <ModalContent top={["0.7em"]} left={[0, 0, "8rem"]} w={["90vh","90vh","420px"]} minH={["550px", "550px" ,"510px"]} borderRadius="15px" boxShadow="2xl">
        <ModalHeader>
          <Flex alignItems="center" justify="space-between">
            <Heading color="gray.700">{ltd.joke.slice(0, 10)}...</Heading>
            <Text color="gray.600" fontSize="3xl" mr="6">
              IT業界
            </Text>
          </Flex>
        </ModalHeader>
        <ModalCloseButton display={["block","none"]}/>
        <ModalBody>
          <Stack divider={<StackDivider borderColor="gray.200" h="full"/>}>
            <StackDivider borderColor="gray.200" />
            <Flex my="2" alignItems="center">
              <Icon as={FcMoneyTransfer} />
              <Text color="gray.600" ml="2">
                830万円
              </Text>
            </Flex>
            <Wrap my="5">
              <Text color="gray.500">{ltd.joke}</Text>
            </Wrap>
            <Wrap my="5">
              <Flex alignItems="center">
                <Icon as={FcInfo} />
                <Text color="gray.700" ml="2">
                  会社情報
                </Text>
              </Flex>
              <Wrap w="full">
                <Tag>家賃補助</Tag>
                <Tag>技術書購入手当</Tag>
                <Tag>海外カンファレンス補助</Tag>
                <Tag>お子さんのバースデー休暇</Tag>
              </Wrap>
            </Wrap>
          </Stack>
        </ModalBody>
        <ModalFooter>
          <Wrap textAlign="center" w="full">
            <Divider/>
              <Center w="full" py="3" onClick={() => alert("report")}>
                <Icon as={FcAdvertising} />
                <Text color="gray.600" ml="2">
                  {ltd.joke.slice(0, 10)}...を報告する
                </Text>
              </Center>
            </Wrap>
        </ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default LtdDetailModal;
