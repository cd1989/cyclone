import { Button, Row, Col, Form } from 'antd';
import { FieldArray, FastField } from 'formik';
import { defaultFormItemLayout } from '@/lib/const';
import Resource from './Form';
import { inject, observer } from 'mobx-react';
import PropTypes from 'prop-types';
import style from '@/components/workflow/component/index.module.less';

const FormItem = Form.Item;
const Fragment = React.Fragment;

@inject('resource')
@observer
class ResourceArray extends React.Component {
  static propTypes = {
    resources: PropTypes.array,
    resourcesField: PropTypes.string,
    type: PropTypes.oneOf(['inputs', 'outputs']),
    update: PropTypes.bool,
    project: PropTypes.string,
    resource: PropTypes.object,
    noLabel: PropTypes.bool,
    resourcesArr: PropTypes.array,
    setFieldValue: PropTypes.func,
    workflowName: PropTypes.string,
  };

  state = {
    visible: false,
  };

  getIntegrationName = _argument => {
    const reg = /^\$\.+/;
    const item = _.find(_argument, o => reg.test(o.value));
    // NOTE: get integration name from $.${namespace}.${integration}/data.integration/sonarQube.server
    if (item) {
      const value = _.get(item, 'value').split('/data.integration');
      const integration = value[0].split('.')[2];
      return integration;
    }
  };

  editResource = (r, index) => {
    const {
      update,
      project,
      resource: { getResource },
    } = this.props;
    let state = {
      modifyData: r,
      visible: true,
      modifyIndex: index,
    };

    if (update) {
      getResource(project, r.name, data => {
        const info = {
          name: _.get(data, 'metadata.name'),
          type: _.get(data, 'spec.type'),
          ..._.pick(data, ['spec.parameters']),
          path: r.path,
          integration: this.getIntegrationName(_.get(data, 'spec.parameters')),
        };
        state.modifyData = info;
        this.setState(state);
      });
    } else {
      this.setState(state);
    }
  };
  render() {
    const {
      resources,
      resourcesField,
      type = 'inputs',
      update,
      project,
      resourcesArr,
      setFieldValue,
      workflowName,
    } = this.props;
    const colSpan = type === 'inputs' ? 6 : 9;
    const { visible, modifyData, modifyIndex } = this.state;
    return (
      <FormItem label={intl.get('sideNav.resource')} {...defaultFormItemLayout}>
        <FieldArray
          name={resourcesField}
          render={arrayHelpers => (
            <div>
              {resources.length > 0 && (
                <Row gutter={16}>
                  <Col span={colSpan}>{intl.get('name')}</Col>
                  <Col span={colSpan}>{intl.get('type')}</Col>
                  {type === 'inputs' && (
                    <Col span={colSpan}>{intl.get('path')}</Col>
                  )}
                </Row>
              )}
              {/* TODO(qme): click resource list show modal and restore resource form */}
              {resources.map((r, i) => {
                return (
                  <FastField
                    key={i}
                    name={`${resourcesField}.${i}`}
                    validate={value =>
                      !_.get(value, 'name')
                        ? { name: intl.get('validate.required') }
                        : undefined
                    }
                    render={({ field, form }) => {
                      const error = _.get(
                        form,
                        `errors.${resourcesField}.${i}.name`
                      );
                      const itemStyle = error
                        ? { border: '1px solid #f5222d', borderRadius: '4px' }
                        : {};
                      return (
                        <Fragment>
                          <div style={itemStyle}>
                            <Row gutter={16}>
                              <Col span={colSpan}>
                                {_.get(r, 'name') || '--'}
                              </Col>
                              <Col span={colSpan}>{_.get(r, 'type')}</Col>
                              {type === 'inputs' && (
                                <Col span={colSpan}>{_.get(r, 'path')}</Col>
                              )}
                              <Col span={6}>
                                <div className={style['resource-action']}>
                                  <Button
                                    type="circle"
                                    icon="edit"
                                    onClick={() => {
                                      this.editResource(r, i);
                                    }}
                                  />
                                  <Button
                                    type="circle"
                                    icon="delete"
                                    onClick={() => arrayHelpers.remove(i)}
                                  />
                                </div>
                              </Col>
                            </Row>
                          </div>
                          {error && (
                            <div style={{ color: '#f5222d' }}>
                              {intl.get('validate.incompleteResource')}
                            </div>
                          )}
                        </Fragment>
                      );
                    }}
                  />
                );
              })}
              <Button
                ico="plus"
                onClick={() => {
                  this.setState({ visible: true });
                }}
              >
                {intl.get('workflow.addResource')}
              </Button>
              {visible && (
                <Resource
                  type={type}
                  handleModalClose={() => {
                    this.setState({ visible: false });
                  }}
                  SetReasourceValue={(value, modify) => {
                    modify
                      ? arrayHelpers.replace(modifyIndex, value)
                      : arrayHelpers.push(value);
                  }}
                  visible={visible}
                  update={update}
                  project={project}
                  modifyData={modifyData}
                  workflowName={workflowName}
                  resourcesArr={resourcesArr}
                  setFieldValue={setFieldValue}
                />
              )}
            </div>
          )}
        />
      </FormItem>
    );
  }
}

export default ResourceArray;
