package definitions

func init() {
	add(`CannotVerifyWithSMP`, &defCannotVerifyWithSMP{})
}

type defCannotVerifyWithSMP struct{}

func (*defCannotVerifyWithSMP) String() string {
	return `<interface>
  <object class="GtkInfoBar" id="error_verifying_smp">
    <property name="message-type">GTK_MESSAGE_WARNING</property>
    <child internal-child="content_area">
      <object class="GtkBox" id="box">
        <property name="homogeneous">false</property>
        <child>
          <object class="GtkLabel" id="message">
            <property name="wrap">true</property>
          </object>
        </child>
      </object>
    </child>
    <child internal-child="action_area">
      <object class="GtkBox" id="button_box">
        <child>
          <object class="GtkButton" id="try_later_button">
            <property name="label" translatable="yes">Sorry</property>
            <property name="can-default">true</property>
          </object>
        </child>
      </object>
    </child>
  </object>
</interface>
`
}