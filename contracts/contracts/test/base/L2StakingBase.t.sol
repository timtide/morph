// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {L2MessageBaseTest} from "./L2MessageBase.t.sol";
import {Predeploys} from "../../libraries/constants/Predeploys.sol";
import {L2Sequencer} from "../../L2/staking/L2Sequencer.sol";
import {Gov} from "../../L2/staking/Gov.sol";
import {Submitter} from "../../L2/submitter/Submitter.sol";
import {Types} from "../../libraries/common/Types.sol";

contract L2StakingBaseTest is L2MessageBaseTest {
    uint256 public beginSeq = 10;
    uint256 public version = 0;
    bytes[] public sequencerBLSKeys;
    address[] public sequencerAddrs;

    // L2Sequencer config
    L2Sequencer public l2Sequencer;

    uint256 public constant SEQUENCER_SIZE = 3;

    // Submitter config
    Submitter public l2Submitter;
    uint256 public NEXT_EPOCH_START = 1700000000;

    // Gov config
    Gov public l2Gov;

    uint256 public PROPOSAL_INTERVAL = 1000;
    uint256 public ROLLUP_EPOCH = 1000;
    uint256 public MAX_CHUNKS = 1000000000;

    function setUp() public virtual override {
        super.setUp();
        // Set the proxy at the correct address
        hevm.etch(
            Predeploys.L2_SEQUENCER,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_SUBMITTER,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        hevm.etch(
            Predeploys.L2_GOV,
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(multisig),
                    new bytes(0)
                )
            ).code
        );
        TransparentUpgradeableProxy l2SequencerProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_SEQUENCER)
            );
        TransparentUpgradeableProxy l2SubmitterProxy = TransparentUpgradeableProxy(
                payable(Predeploys.L2_SUBMITTER)
            );
        TransparentUpgradeableProxy l2GovProxy = TransparentUpgradeableProxy(
            payable(Predeploys.L2_GOV)
        );
        hevm.store(
            address(l2SequencerProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2SubmitterProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );
        hevm.store(
            address(l2GovProxy),
            bytes32(PROXY_OWNER_KEY),
            bytes32(abi.encode(address(multisig)))
        );

        hevm.startPrank(multisig);
        // deploy impl contracts
        L2Sequencer l2SequencerImpl = new L2Sequencer(
            payable(NON_ZERO_ADDRESS)
        );
        Gov govImpl = new Gov();
        Submitter submitterImpl = new Submitter();

        // upgrade proxy
        Types.SequencerInfo[] memory sequencerInfos = new Types.SequencerInfo[](
            SEQUENCER_SIZE
        );
        for (uint256 i = 0; i < SEQUENCER_SIZE; i++) {
            address user = address(uint160(beginSeq + i));
            Types.SequencerInfo memory sequencerInfo = ffi.generateStakingInfo(
                user
            );
            sequencerInfos[i] = sequencerInfo;
            sequencerAddrs.push(sequencerInfo.addr);
        }
        ITransparentUpgradeableProxy(address(l2SequencerProxy))
            .upgradeToAndCall(
                address(l2SequencerImpl),
                abi.encodeWithSelector(
                    L2Sequencer.initialize.selector,
                    sequencerInfos
                )
            );
        ITransparentUpgradeableProxy(address(l2GovProxy)).upgradeToAndCall(
            address(govImpl),
            abi.encodeWithSelector(
                Gov.initialize.selector,
                PROPOSAL_INTERVAL, // _proposalInterval
                0, // _batchBlockInterval
                0, // _batchMaxBytes
                FINALIZATION_PERIOD_SECONDS, // _batchTimeout
                ROLLUP_EPOCH, // rollupEpoch
                MAX_CHUNKS // maxChunks
            )
        );
        ITransparentUpgradeableProxy(address(l2SubmitterProxy))
            .upgradeToAndCall(
                address(submitterImpl),
                abi.encodeWithSelector(
                    Submitter.initialize.selector,
                    sequencerAddrs,
                    block.timestamp
                )
            );

        // set address
        l2Sequencer = L2Sequencer(payable(address(l2SequencerProxy)));
        l2Gov = Gov(payable(address(l2GovProxy)));
        l2Submitter = Submitter(payable(address(l2SubmitterProxy)));

        _changeAdmin(address(l2Sequencer));
        _changeAdmin(address(l2Gov));
        _changeAdmin(address(l2Submitter));

        hevm.stopPrank();
    }
}
