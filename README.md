# XuperChain Crypto

This project is the crypto library for XuperChain.

����Ŀ�ǳ�������ص������ģ�顣Ŀǰ��Դ�����Ѿ�֧����Xuperͳһ����ǩ���㷨��Schnorrǩ������ǩ��������ǩ���ȶ���ǩ���㷨��Ҳ֧�������ܷ����ֲ�ȷ���Լӽ��ܵȶ�������ѧ������

NISTϵ���㷨�����ã�
"github.com/xuperchain/crypto/client/service/xchain"

��ϸ����˵����ο�����ĺ���ע��

ʹ�����ӣ�
```
import (
	"log"

	"github.com/xuperchain/crypto/client/service/xchain"
	"github.com/xuperchain/crypto/core/account"
	"github.com/xuperchain/crypto/core/hdwallet/rand"
)

xcc := new(xchain.XchainCryptoClient)

ecdsaAccount, err := xcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
if err != nil {
	log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
	return
}

log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
```

------

����ϵ���㷨�����ã�
"github.com/xuperchain/crypto/client/service/gm"

��ϸ����˵����ο�����ĺ���ע��

ʹ�����ӣ�
```
import (
	"log"

	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/hdwallet/rand"
)

gcc := new(gm.GmCryptoClient)

ecdsaAccount, err := gcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
if err != nil {
	log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
	return
}
log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
```