
一直想实现一下B+树,正好最近在学golang所以就用golang来实现了.
具体原理参考:https://yq.aliyun.com/articles/9280

添加记录的测试代码在 BTreePlus_test.go里的TestBTreePlus_Insert()

    根据10000个大随机数来生成10000个记录然后加入到B+树中.
    其中每添加一次记录都调用testAllNature()来验证是否还满足b+树的性质


删除记录的测试代码在 BTreePlus_test.go里TestBTreePlus_Remove()里

    根据10000个大随机数来生成10000个记录然后加入到B+树中.
    然后再删除这10000个记录
    其中每删除一次记录都调用testAllNature()来验证是否还满足b+树的性质

testAllNature()通过测试B+树的6个性质来验证所构造的树是否是B+树





