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
          <Stack divider={<StackDivider borderColor="gray.200" />}>
            <StackDivider borderColor="gray.200" />
            <Flex my="2" alignItems="center">
              <Icon as={FcMoneyTransfer} />
              <Text color="gray.600" ml="2">
                830万円
              </Text>
            </Flex>
            <Wrap my="5">
              <Text color="gray.500">
                (会社紹介)Lorem ipsum is placeholder text commonly used in the
                graphic, print, and publishing industries for previewing layouts
                and visual mockups.
              </Text>
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
            <Wrap textAlign="center" w="full">
              <Center w="full" onClick={() => alert("report")}>
                <Icon as={FcAdvertising} />
                <Text color="gray.600" ml="2">
                  {ltd.name}を報告する
                </Text>
              </Center>
            </Wrap>
          </Stack>
        </ModalBody>
        <ModalFooter></ModalFooter>
      </ModalContent>
    </Modal>
  );
};

export default LtdDetailModal;
