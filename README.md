
一直想实现一下B+树,正好最近在学golang所以就用golang来实现了.
具体原理参考:https://yq.aliyun.com/articles/9280

测试代码在BTreePlus_test.go里
根据100000个大随机数来生成100000个记录然后加入到B+树中.

通过测试B+树的6个性质来验证所构造的B+树是否正确
