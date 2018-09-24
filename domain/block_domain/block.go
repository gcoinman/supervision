package block_domain

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/D-Technologies/supervision/domain/eth_domain"
	"github.com/D-Technologies/supervision/domain/received_tx_domain"
)

// Block represents a ethereum block in domain layer
type Block struct {
	Transactions []*eth_domain.Tx
}

// Scan scans all transactions in a block and see if transactions with specified conditions exist
func (b *Block) Scan(contractAddr string, receiveAddrs []string) []*received_tx_domain.ReceivedTx {
	var matchedTxs []*received_tx_domain.ReceivedTx
	for _, tx := range b.Transactions {
		// tx.To address must be the same as contract address
		// transfer(from, to, tokenID) or safeTransfer(from, to, tokenID)
		// 0x(8bit), method signature(32bit), first parameter(256bit), second parameter(256bit), third parameter(256bit)
		if strings.ToLower(tx.To) != strings.ToLower(contractAddr) || len(tx.Data) != 202 {
			break
		}

		// If method signature is not equal to sha3 of
		// transfer(addressaddress,uint256)
		// safeTransfer(address,address,uint256)
		// safeTransfer(address,address,uint256,bytes)
		// then breaks
		if methodSig := strings.ToLower(tx.Data[2:10]); methodSig != "23b872dd" && methodSig != "42842e0e" && methodSig != "b88d4fde" {
			break
		}

		for _, addr := range receiveAddrs {
			// If to-address is not equal to receive address (which you specify) then breaks
			if to := fmt.Sprintf("0x%s", removeZeros(tx.Data[75:138])); strings.ToLower(to) != strings.ToLower(addr) {
				break
			}

			blockNum, _ := strconv.ParseInt(tx.BlockNum, 0, 64)
			tokenID, _ := strconv.ParseInt(removeZeros(tx.Data[139:202]), 16, 64)
			from := fmt.Sprintf("0x%s", removeZeros(tx.Data[11:74]))

			rt := &received_tx_domain.ReceivedTx{
				Hash:         tx.Hash,
				BlockNum:     blockNum,
				From:         from,
				ReceiveAddir: addr,
				TokenID:      tokenID,
				Status:       received_tx_domain.Pending,
			}

			fmt.Printf("\nDeposit was detected at %d, from %s to %s. TokenID is %d.\n\n", rt.BlockNum, rt.From, addr, rt.TokenID)

			matchedTxs = append(matchedTxs, rt)
		}
	}

	return matchedTxs
}

func removeZeros(str string) string {
	rep := regexp.MustCompile(`^0+`)
	return rep.ReplaceAllString(str, "")
}
