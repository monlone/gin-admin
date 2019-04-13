import React, { PureComponent } from 'react';
import { connect } from 'dva';
import { Card, Divider } from 'antd';
import MerchantCard from './MerchantCard';
import DescriptionList from '@/components/DescriptionList';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';

const { Description } = DescriptionList;

@connect(state => ({
  loading: state.loading.models.merchant,
  merchant: state.merchant,
}))
class MerchantDetail extends PureComponent {
  constructor(props) {
    super(props);

    const { dispatch } = this.props;
    const params = props.location.query;
    dispatch({
      type: 'merchant/fetchForm',
      payload: {
        record_id: params.record_id,
      },
    });
  }

  componentDidMount() {}

  dispatch = action => {
    const { dispatch } = this.props;
    dispatch(action);
  };

  renderDataForm() {
    return <MerchantCard onCancel={this.onDataFormCancel} onSubmit={this.onDataFormSubmit} />;
  }

  render() {
    const {
      loading,
      merchant: { formData },
    } = this.props;

    const breadcrumbList = [{ title: '商户管理' }, { title: '商户详情', href: '/merchant/detail' }];

    return (
      <PageHeaderWrapper title="商户管理" breadcrumbList={breadcrumbList} loading={loading}>
        <Card bordered={false}>
          <DescriptionList size="large" title="商户详情" style={{ marginBottom: 32 }}>
            <Description term="商户号">{formData && formData.record_id}</Description>
            <Description term="状态">{formData && formData.status}</Description>
          </DescriptionList>
          <Divider style={{ marginBottom: 32 }} />
        </Card>
      </PageHeaderWrapper>
    );
  }
}

export default MerchantDetail;
